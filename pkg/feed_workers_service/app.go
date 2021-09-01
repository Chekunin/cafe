package feed_workers_service

import (
	"cafe/pkg/common"
	"cafe/pkg/common/catcherr"
	log "cafe/pkg/common/logman"
	_ "cafe/pkg/common/logman/drivers/stack"
	_ "cafe/pkg/common/logman/drivers/zap"
	"cafe/pkg/common/utils"
	httpDbManaber "cafe/pkg/db_manager/http"
	"cafe/pkg/feed_queue/workers"
	"context"
	"errors"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/getsentry/sentry-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Config struct {
	Listen         string                      `yaml:"listen"`
	DocPath        string                      `yaml:"docPath"`
	SentryDSN      string                      `yaml:"sentry_dsn"`
	SentryENV      string                      `yaml:"sentry_env"`
	Logger         log.Config                  `yaml:"logger"`
	LoggerChannels log.ChannelArbitraryConfigs `yaml:"logChannels"`
	DbManagerURL   string                      `yaml:"db_manager_url"`
	Redis          RedisConfig                 `yaml:"redis"`
}

type DbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	PoolSize int    `yaml:"pool_size"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type App struct {
	config      Config
	workersPool *workers.Workers
	route       *gin.Engine
	doneChan    chan bool
}

func NewApp(config Config) *App {
	log.InitOrPanic(config.Logger.WithChannels(config.LoggerChannels))

	catcherr.InitOrPanic(catcherr.Config{
		Sentry: sentry.ClientOptions{
			Dsn:         config.SentryDSN,
			Environment: config.SentryENV,
		},
	}.WithLogger(log.Current()))

	if log.Current().Level() == log.DebugLevel {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(common.RequestIdMiddleware())
	r.Use(common.RequestLogger())
	r.Use(common.ErrorLogger())
	r.Use(common.ErrorResponder(func(c *gin.Context, code int, obj interface{}) {
		c.Data(0, "application/x-gob", utils.ToGobBytes(obj))
	}))
	r.Use(common.Recovery())

	r.Use(cors.New(utils.GetCorsConfigs()))

	r.NoRoute(func(c *gin.Context) {
		c.AbortWithError(http.StatusNotFound, common.ErrPageNotFound)
	})

	r.GET("/health", healthCheck())
	// временное решение. Так мы сразу будет отдавать положительный ответ на /health,
	// но не будет обрабатывать остальные запросы
	srv2 := &http.Server{
		Addr:    config.Listen,
		Handler: r,
	}
	go func() {
		log.Info("Start listening for health-check", log.Fields{"port": config.Listen})
		if err := srv2.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			catcherr.AsCritical().Catch(wrapErr.NewWrapErr(fmt.Errorf("ListenAndServe for health-check"), err))
		}
	}()
	defer func() {
		if err := srv2.Shutdown(context.TODO()); err != nil {
			catcherr.Catch(wrapErr.NewWrapErr(fmt.Errorf("srv2 server forced to shutdown"), err))
		}
	}()
	defer srv2.Close()

	if config.DocPath != "" {
		r.Static("/doc/api", config.DocPath)
	} else {
		log.Warning("Document path is undefined")
	}

	dbManager, err := httpDbManaber.NewHttpDbManager(config.DbManagerURL)
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("NewHttpDbManager"), err)
		catcherr.AsCritical().CatchAndExit(err)
	}

	workersPool := workers.NewWorkers(workers.NewWorkersParams{
		RedisClientOpt: asynq.RedisClientOpt{
			Addr: fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port),
			//Username: config.Redis.User,
			//Password: config.Redis.Password,
			//DB:       config.Redis.DB,
		},
		DbManager: dbManager,
	})

	app := App{
		config:      config,
		workersPool: workersPool,
		route:       r,
		doneChan:    make(chan bool, 1),
	}

	return &app
}

func (a *App) Run() {
	a.route.GET("/readiness", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	srv := &http.Server{
		Addr:    a.config.Listen,
		Handler: a.route,
	}

	go func() {
		log.Info("Start listening", log.Fields{"port": a.config.Listen})
		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Info(fmt.Sprintf("listen: %s\n", err))
		} else if err != nil {
			catcherr.AsCritical().Catch(wrapErr.NewWrapErr(fmt.Errorf("srv ListenAndServe"), err))
		}
	}()

	if err := a.workersPool.Run(); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("workersPool Run"), err)
		catcherr.AsCritical().Catch(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
	case <-a.doneChan:
	}
	log.Info("Shutdown Server (timeout of 3 seconds) ...")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		catcherr.Catch(wrapErr.NewWrapErr(fmt.Errorf("Server forced to shutdown"), err))
	}

	a.workersPool.Shutdown()

	<-ctx.Done()
	log.Info("Server exiting")
}

func (a *App) Close() {
	a.doneChan <- true
}

func healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Ok")
	}
}
