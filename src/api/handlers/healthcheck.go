package handlers

import (
	"net/http"

	"github.com/allanassis/reddere/src/api/errors"
	"github.com/allanassis/reddere/src/api/response"
	"github.com/allanassis/reddere/src/observability/logging"
	"github.com/allanassis/reddere/src/storages"
	"github.com/labstack/echo/v4"
)

func Healthcheck(db storages.Storage, logger *logging.Logger) func(c echo.Context) error {
	return func(c echo.Context) error {
		err := db.Healthcheck()
		if err != nil {
			logger.Error(string(errors.API_HEALTHCHECK_ERROR), logging.String("error", err.Error()))

			resp := response.NewApiResponse(
				response.WithEventID(c.Get("eventID").(string)),
				response.WithError(errors.API_HEALTHCHECK_ERROR),
			)

			return c.JSON(http.StatusInternalServerError, resp)
		}

		resp := response.NewApiResponse(
			response.WithEventID(c.Get("eventID").(string)),
			response.WithMessage("Working"),
		)
		return c.JSON(http.StatusOK, resp)
	}
}
