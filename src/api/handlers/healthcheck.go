package handlers

import (
	"net/http"

	"github.com/allanassis/reddere/src/observability"
	"github.com/allanassis/reddere/src/storages"
	"github.com/labstack/echo/v4"
)

func Healthcheck(db storages.Storage, logger *observability.Logger) func(c echo.Context) error {
	return func(c echo.Context) error {
		err := db.Healthcheck()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Working"})
	}
}
