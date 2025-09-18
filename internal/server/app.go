package server

import (
	"context"
	"item-comparison-ai-api/config"
	"item-comparison-ai-api/internal/database"
	h "item-comparison-ai-api/internal/handlers"
	"item-comparison-ai-api/internal/logger"
	middlewares "item-comparison-ai-api/internal/middleware"
	"net/http"

	ginlogrus "github.com/toorop/gin-logrus"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Application - represents a application server configuration
type Application struct {
	config     *config.AppConfig
	database   *database.Database
	httpServer *http.Server
	router     *gin.Engine
	logger     logger.Logger
}

// New - responsible to creates a new instance from Application
func New(config *config.AppConfig, db *database.Database, router *gin.Engine, logger logger.Logger) *Application {
	gin.SetMode(getGinExecMode(config))

	var server = &http.Server{
		Addr:    config.BindAddr,
		Handler: router,
	}

	var application = &Application{
		config:     config,
		database:   db,
		httpServer: server,
		router:     router,
		logger:     logger,
	}

	return application
}

func getGinExecMode(c *config.AppConfig) string {
	if c.Environment == "local" {
		return gin.DebugMode
	}

	return gin.ReleaseMode
}

// WithMiddlewares - responsible to attach middlewares into http request pipeline
func (a *Application) WithMiddlewares() *Application {
	a.router.Use(cors.Default())
	a.router.Use(ginlogrus.Logger(a.logger.GetLogger()))
	a.router.NoRoute(func(ctx *gin.Context) {
		h.HandleError(ctx, h.ErrNotFound)
	})

	a.router.Use(gin.Recovery())
	return a
}

// WithHandlers ...
func (a *Application) WithHandlers(routePrefix string, handlers ...Bindable) *Application {
	var router = a.router.Group(routePrefix)

	for _, handler := range handlers {
		handler.Bind(router, a)
	}

	return a
}

// Start ...
func (a *Application) Start() {
	var logger = a.logger.GetLogger()

	a.httpServer.Handler = a.router
	if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("listen: %s\n", err)
	}
}

// Shutdown ...
func (a *Application) Shutdown(ctx context.Context) error {
	return a.Shutdown(ctx)
}

// WithHealthcheck ...
func (a *Application) WithHealthcheck() *Application {
	a.router.GET("/health", middlewares.Health(a.database, a.config))
	return a
}
