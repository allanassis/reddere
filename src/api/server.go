package api

import (
	"github.com/allanassis/reddere/src/api/handlers"
	"github.com/allanassis/reddere/src/storages"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitServer(db storages.Storage) {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/healthcheck", handlers.Healthcheck(db))

	// Template
	e.POST("/template", handlers.PostTemplate(db))
	e.GET("/template/:id", handlers.GetTemplate(db))

	e.Logger.Fatal(e.Start(":1323"))
}
