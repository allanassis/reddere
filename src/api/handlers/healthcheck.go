package handlers

import (
	"net/http"

	"github.com/allanassis/reddere/src/observability/logging"
	"github.com/allanassis/reddere/src/storages"
	"github.com/labstack/echo/v4"
)

func Healthcheck(db storages.Storage, logger *logging.Logger) func(c echo.Context) error {
	return func(c echo.Context) error {
		err := db.Healthcheck()
		if err != nil {
			logger.Error(string(API_HEALTHCHECK_ERROR), logging.String("error", err.Error()))
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{
					"errorCode": API_HEALTHCHECK_ERROR.String(),
					"message":   string(API_HEALTHCHECK_ERROR),
					"eventID":   c.Get("eventID").(string),
				},
			)
		}
		return c.JSON(http.StatusOK, map[string]string{"message": "Working", "eventID": c.Get("eventID").(string)})
	}
}
