package db_manager_service

import (
	"cafe/pkg/common"
	"cafe/pkg/common/catcherr"
	log "cafe/pkg/common/logman"
	_ "cafe/pkg/common/logman/drivers/stack"
	_ "cafe/pkg/common/logman/drivers/zap"
	"cafe/pkg/common/utils"
	"cafe/pkg/db_manager/memory"
	"cafe/pkg/db_manager_service/api/delivery/rest"
	"cafe/pkg/db_manager_service/api/usecase"
	"context"
	"errors"
	"fmt"
	wrapErr "github.com/Chekunin/wraperr"
	"github.com/getsentry/sentry-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
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
	Db             *DbConfig                   `yaml:"db"`
	Logger         log.Config                  `yaml:"logger"`
	LoggerChannels log.ChannelArbitraryConfigs `yaml:"logChannels"`
}

type DbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	PoolSize int    `yaml:"pool_size"`
}

type App struct {
	config   Config
	usecase  *usecase.Usecase
	route    *gin.Engine
	doneChan chan bool
	db       *pg.DB
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

	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Db.Host, config.Db.Port),
		User:     config.Db.User,
		Password: config.Db.Password,
		Database: config.Db.Name,
		PoolSize: config.Db.PoolSize,
	})

	dbManager, err := memory.NewDbManager(memory.NewDbManagerParams{DB: db})
	if err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("NewDbManager"), err)
		catcherr.AsCritical().CatchAndExit(err)
	}

	usecase := usecase.NewUsecase(dbManager)
	rest.NewRest(r.Group("/v1"), usecase)

	app := App{
		config:   config,
		route:    r,
		usecase:  usecase,
		doneChan: make(chan bool, 1),
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

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-quit:
	case <-a.doneChan:
	}
	log.Info("Shutdown Server (timeout of 3 seconds) ...")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := a.db.Close(); err != nil {
		catcherr.Catch(wrapErr.NewWrapErr(fmt.Errorf("close db"), err))
	}

	if err := srv.Shutdown(ctx); err != nil {
		catcherr.Catch(wrapErr.NewWrapErr(fmt.Errorf("Server forced to shutdown"), err))
	}

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
