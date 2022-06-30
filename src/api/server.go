package api

import (
	"github.com/allanassis/reddere/src/api/handlers"
	"github.com/allanassis/reddere/src/observability"
	"github.com/allanassis/reddere/src/storages"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitServer(db storages.Storage, logger *observability.Logger) {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/healthcheck", handlers.Healthcheck(db, logger))

	// Template
	e.POST("/template", handlers.PostTemplate(db, logger))
	e.GET("/template/:id", handlers.GetTemplate(db, logger))
	e.DELETE("/template/:id", handlers.DeleteTemplate(db, logger))

	e.Logger.Fatal(e.Start(":1323"))
}
