package api

import (
	"github.com/allanassis/reddere/src/api/handlers"
	customMiddlewares "github.com/allanassis/reddere/src/api/middlewares"
	"github.com/allanassis/reddere/src/config"
	"github.com/allanassis/reddere/src/observability/logging"
	"github.com/allanassis/reddere/src/services"
	"github.com/allanassis/reddere/src/storages"

	_ "github.com/allanassis/reddere/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Reddere
// @version 1.0
// @description This is a simple server to save templates.
// @contact.name Allan Assis de Andrade
// @license.name MIT License
// @BasePath /
func NewServer(service services.Service, config *config.Config, db storages.Storage, logger *logging.Logger) {
	e := echo.New()

	// Middleware pre routes execution
	e.Pre(customMiddlewares.RequestLogger(logger))

	// Middleware pos routes execution
	e.Use(middleware.Recover())
	e.Use(customMiddlewares.ResponseLogger(logger))

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/healthcheck", handlers.Healthcheck(db, logger))

	// Template
	e.POST("/template", handlers.PostTemplate(service, db, logger))
	e.GET("/template/:id", handlers.GetTemplate(service, db, logger))
	e.DELETE("/template/:id", handlers.DeleteTemplate(service, db, logger))

	e.Logger.Fatal(e.Start(":1323"))
}
