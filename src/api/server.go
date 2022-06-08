package api

import (
	"github.com/allanassis/reddere/src/api/handlers"
	"github.com/allanassis/reddere/src/storages"

	"github.com/labstack/echo/v4"
)

func InitServer(db storages.Storage) {
	e := echo.New()

	e.GET("/healthcheck", handlers.Healthcheck(db))
	e.Logger.Fatal(e.Start(":1323"))
}
