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

		loggingFields := []logging.Field{logging.String("entity", "Template")}
		err := c.Bind(template)
		if err != nil {
			logger.Error("Error when binding payload",
				append(loggingFields, logging.String("error", err.Error()))...,
			)
			c.JSON(
				http.StatusUnprocessableEntity,
				map[string]interface{}{"msg": "error when binding payload, you payload is invalid"},
			)
		}
		loggingFields = append(loggingFields, logging.Any("template", template))
		logger.Debug("Received request with payload", loggingFields...)

		id, err := service.Save(template)
		if err != nil {
			logger.Error("Internal Error when saving template",
				append(loggingFields, logging.String("error", err.Error()))...,
			)
			c.JSON(
				http.StatusInternalServerError,
				map[string]interface{}{"msg": "Internal Error when saving template"},
			)
		}
		logger.Info("Succefuly saved template", append(loggingFields, logging.String("id", id))...)

		template.ID = id

		return c.JSON(http.StatusCreated, template)
	}
}

func GetTemplate(service services.Service, storage storages.Storage, logger *logging.Logger) func(c echo.Context) error {
	return func(c echo.Context) error {
		template := new(entities.Template)

		loggingFields := []logging.Field{logging.String("entity", "Template")}

		templateId := c.Param("id")
		err := service.Build(template, templateId)
		if err != nil {
			logger.Error("Error when binding template",
				append(loggingFields, logging.String("error", err.Error()))...,
			)
			c.JSON(
				http.StatusInternalServerError,
				map[string]interface{}{"msg": "error when binding template"},
			)
		}
		loggingFields = append(loggingFields, logging.Any("template", template))
		logger.Info("Succefuly get template", append(loggingFields, logging.String("id", templateId))...)

		return c.JSON(http.StatusOK, template)
	}
}

func DeleteTemplate(service services.Service, storage storages.Storage, logger *logging.Logger) func(c echo.Context) error {
	return func(c echo.Context) error {
		loggingFields := []logging.Field{logging.String("entity", "Template")}

		templateId := c.Param("id")

		err := service.Delete(templateId, "template")
		if err != nil {
			logger.Error("Error when binding template",
				append(loggingFields, logging.String("error", err.Error()))...,
			)
			c.JSON(
				http.StatusInternalServerError,
				map[string]interface{}{"msg": "error when binding template"},
			)
		}
		logger.Info("Succefuly deleted template", append(loggingFields, logging.String("id", templateId))...)

		return c.JSON(http.StatusOK, templateId)
	}
}
