package catcherr

import (
	"cafe/pkg/common/logman"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
)

var (
	MultipleInitErr = errors.New("Catcherr already initialized")

	toLogmanLevel = map[errLevel]logman.Level{
		criticalLevel: logman.CriticalLevel,
		errorLevel:    logman.ErrorLevel,
		warningLevel:  logman.WarningLevel,
	}
)

type errLevel uint8

const (
	criticalLevel errLevel = iota + 1
	errorLevel
	warningLevel

	defaultLevel errLevel = errorLevel
)

type Config struct {
	Logger               logman.Logger
	Sentry               sentry.ClientOptions
	DisableErrorTracking bool

	errorTracker       string
	errorTrackingLevel errLevel
}

func (c Config) WithLogger(l logman.Logger) Config {
	c.Logger = l

	return c
}

var catcher = newDefault()

func AsWarning() errRecord {
	return catcher.asLevel(warningLevel)
}
func AsError() errRecord {
	return catcher.asLevel(errorLevel)
}
func AsCritical() errRecord {
	return catcher.asLevel(criticalLevel)
}
func NotTrackable() errRecord {
	return catcher.trackable(false)
}
func Trackable() errRecord {
	return catcher.trackable(true)
}
func Catch(err error, fields ...Fields) {
	catcher.asLevel(defaultLevel).Catch(err, fields...)
}
func CatchAndExit(err error, fields ...Fields) {
	catcher.asLevel(defaultLevel).CatchAndExit(err, fields...)
}

type catcherr struct {
	errorTracker           string
	errorTrackingLevel     errLevel
	isErrorTrackingEnabled bool
	isInited               bool
	logger                 logman.Logger
}
type Fields map[string]interface{}

func Init(cfg Config) error {
	if catcher.isInited {
		return fmt.Errorf("Catcherr init failed <= %w", MultipleInitErr)
	}

	c, err := New(cfg)
	if err != nil {
		return fmt.Errorf("Catcherr init failed <= %w", err)
	}

	catcher = c
	catcher.isInited = true

	return nil
}
func InitOrPanic(cfg Config) {
	if err := Init(cfg); err != nil {
		panic(err)
	}
}
func New(cfg Config) (*catcherr, error) {
	c := &catcherr{
		isErrorTrackingEnabled: !cfg.DisableErrorTracking,
		logger:                 cfg.Logger,
	}

	if c.isErrorTrackingEnabled {
		if err := c.initErrorTracker(cfg); err != nil {
			return nil, fmt.Errorf("Error tracker init failed <= %w", err)
		}
	}

	c.logger.Debug("Catcherr created")

	return c, nil
}
func NewOrPanic(cfg Config) *catcherr {
	c, err := New(cfg)
	if err != nil {
		panic(err)
	}

	return c
}
func newDefault() *catcherr {
	return NewOrPanic(Config{Logger: logman.Current()})
}
func Close() {
	catcher.close()
}

func (c *catcherr) initErrorTracker(cfg Config) error {
	c.errorTracker = cfg.errorTracker
	if c.errorTracker == "" {
		c.errorTracker = "sentry"
	}

	c.errorTrackingLevel = cfg.errorTrackingLevel
	if c.errorTrackingLevel == 0 {
		c.errorTrackingLevel = errorLevel
	}

	switch c.errorTracker {
	case "sentry":
		if err := c.initSentry(cfg.Sentry); err != nil {
			return fmt.Errorf("Sentry init failed <= %w", err)
		}
	default:
		return fmt.Errorf("Unknown error tracker \"%s\"", c.errorTracker)
	}

	c.logger.Debug(
		"Error tracker client is initialized",
		logman.Fields{"tracker": c.errorTracker},
	)

	return nil
}
func (c *catcherr) initSentry(cfg sentry.ClientOptions) error {
	if err := sentry.Init(cfg); err != nil {
		return err
	}

	return nil
}
func (c *catcherr) asLevel(l errLevel) errRecord {
	switch defaultLevel {
	case criticalLevel, errorLevel, warningLevel:
		return errRecord{catcherr: c, level: l}
	default:
		catcher.catch(catcher.asLevel(errorLevel).
			withError(errors.New("Unknown error level"), Fields{"level": l}))
		return errRecord{catcherr: c, level: errorLevel}
	}
}
func (c *catcherr) trackable(val bool) errRecord {
	return errRecord{
		catcherr:    c,
		level:       defaultLevel,
		isTrackable: isTrackable{isSet: true, val: val},
	}
}
func (c *catcherr) catch(e errRecord) {
	c.logger.Log(
		toLogmanLevel[e.level],
		e.error.Error(),
		toLogmanFields(e.fields),
	)

	if !c.isErrorTrackingEnabled {
		return
	}

	if e.isTrackable.isSet && !e.isTrackable.val {
		return
	}

	if !e.isTrackable.isSet && c.errorTrackingLevel < e.level {
		return
	}

	switch c.errorTracker {
	case "sentry":
		if len(e.fields) > 0 {
			sentry.ConfigureScope(func(scope *sentry.Scope) {
				scope.SetContext("Fields", e.fields)
			})
		}
		sentry.CaptureException(e.error)
	default:
		c.logger.Log(
			toLogmanLevel[errorLevel],
			"Unknown error tracker",
			toLogmanFields(Fields{"tracker": c.errorTracker}),
		)
	}
}
func (c *catcherr) catchAndExit(e errRecord) {
	c.catch(e)
	c.close()
	os.Exit(1)
}
func (c *catcherr) close() {
	if c.isErrorTrackingEnabled {
		if c.errorTracker == "sentry" {
			sentry.Flush(3 * time.Second)
		}
	}

	c.logger.Debug("Catcherr closed")
	c = nil
}

func toLogmanFields(fields ...Fields) logman.Fields {
	errFields := Fields{}
	for _, fieldSet := range fields {
		for f, v := range fieldSet {
			errFields[f] = v
		}
	}

	return logman.Fields{"error": errFields}
}

type errRecord struct {
	catcherr    *catcherr
	error       error
	fields      Fields
	level       errLevel
	isTrackable isTrackable
}
type isTrackable struct {
	isSet bool
	val   bool
}

func (e errRecord) withError(err error, f ...Fields) errRecord {
	e.error = err

	if len(f) > 0 && e.fields == nil {
		e.fields = Fields{}
	}

	for _, fields := range f {
		for f, v := range fields {
			e.fields[f] = v
		}
	}

	return e
}
func (e errRecord) NotTrackable() errRecord {
	e.isTrackable = isTrackable{isSet: true, val: false}

	return e
}
func (e errRecord) Trackable() errRecord {
	e.isTrackable = isTrackable{isSet: true, val: true}

	return e
}
func (e errRecord) Catch(err error, fields ...Fields) {
	catcher.catch(e.withError(err, fields...))
}
func (e errRecord) CatchAndExit(err error, fields ...Fields) {
	catcher.catchAndExit(e.withError(err, fields...))
}
