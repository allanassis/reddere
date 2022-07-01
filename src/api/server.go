package api

import (
	"github.com/allanassis/reddere/src/api/handlers"
	"github.com/allanassis/reddere/src/observability/logging"
	"github.com/allanassis/reddere/src/services"
	"github.com/allanassis/reddere/src/storages"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitServer(service services.Service, db storages.Storage, logger *logging.Logger) {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/healthcheck", handlers.Healthcheck(db, logger))

	// Template
	e.POST("/template", handlers.PostTemplate(service, db, logger))
	e.GET("/template/:id", handlers.GetTemplate(service, db, logger))
	e.DELETE("/template/:id", handlers.DeleteTemplate(service, db, logger))

	e.Logger.Fatal(e.Start(":1323"))
}
