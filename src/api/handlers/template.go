package handlers

import (
	"net/http"

	"github.com/allanassis/reddere/src/observability/logging"
	"github.com/allanassis/reddere/src/services"
	"github.com/allanassis/reddere/src/services/entities"
	"github.com/allanassis/reddere/src/storages"
	"github.com/labstack/echo/v4"
)

func PostTemplate(service services.Service, storage storages.Storage, logger *logging.Logger) func(c echo.Context) error {
	return func(c echo.Context) error {
		template := new(entities.Template)

		err := c.Bind(template)
		if err != nil {
			panic(err)
		}

		id, err := service.Save(template)
		if err != nil {
			panic(err)
		}

		template.ID = id

		return c.JSON(http.StatusCreated, template)
	}
}

func GetTemplate(service services.Service, storage storages.Storage, logger *logging.Logger) func(c echo.Context) error {
	return func(c echo.Context) error {
		template := new(entities.Template)

		templateId := c.Param("id")
		err := service.Build(template, templateId)
		if err != nil {
			panic(err)
		}

		return c.JSON(http.StatusOK, template)
	}
}

func DeleteTemplate(service services.Service, storage storages.Storage, logger *logging.Logger) func(c echo.Context) error {
	return func(c echo.Context) error {
		templateId := c.Param("id")

		err := service.Delete(templateId)
		if err != nil {
			panic(err)
		}

		return c.JSON(http.StatusOK, templateId)
	}
}
