package handlers

import (
	"net/http"

	"github.com/allanassis/reddere/src/storages"
	"github.com/labstack/echo/v4"
)

func Healthcheck(db storages.Storage) func(c echo.Context) error {
	return func(c echo.Context) error {
		err := db.Healthcheck()
		if err != nil {
			panic(err)
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Working"})
	}
}
