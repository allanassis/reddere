package handlers

import (
	"net/http"

	"github.com/allanassis/reddere/src/services"
	"github.com/allanassis/reddere/src/storages"
	"github.com/labstack/echo/v4"
)

func PostTemplate(storage storages.Storage) func(c echo.Context) error {
	return func(c echo.Context) error {
		template := new(services.Template)

		err := c.Bind(template)
		if err != nil {
			panic(err)
		}

		id, err := template.Save(storage)
		if err != nil {
			panic(err)
		}

		template.ID = id

		return c.JSON(http.StatusCreated, template)
	}
}