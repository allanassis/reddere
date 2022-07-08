package handlers

import (
	"net/http"

	"github.com/allanassis/reddere/src/api/errors"
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
			logger.Error(string(errors.API_BIND_PAYLOAD_ERROR),
				append(loggingFields, logging.String("error", err.Error()))...,
			)
			return c.JSON(
				http.StatusUnprocessableEntity,
				map[string]string{
					"errorCode": errors.API_BIND_PAYLOAD_ERROR.String(),
					"message":   string(errors.API_BIND_PAYLOAD_ERROR),
					"eventID":   c.Get("eventID").(string),
				})
		}
		loggingFields = append(loggingFields, logging.Any("template", template))
		logger.Debug("Received request with payload", loggingFields...)

		id, err := service.Save(template)
		if err != nil {
			logger.Error(string(errors.API_POST_ERROR),
				append(loggingFields, logging.String("error", err.Error()))...,
			)
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{
					"errorCode": errors.API_POST_ERROR.String(),
					"message":   string(errors.API_POST_ERROR),
					"eventID":   c.Get("eventID").(string),
				})
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
			logger.Error(string(errors.API_GET_ERROR),
				append(loggingFields, logging.String("error", err.Error()))...,
			)
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{
					"errorCode": errors.API_GET_ERROR.String(),
					"message":   string(errors.API_GET_ERROR),
					"eventID":   c.Get("eventID").(string),
				})
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
			logger.Error(string(errors.API_DELETE_ERROR),
				append(loggingFields, logging.String("error", err.Error()))...,
			)
			return c.JSON(
				http.StatusInternalServerError,
				map[string]string{
					"errorCode": errors.API_DELETE_ERROR.String(),
					"message":   string(errors.API_DELETE_ERROR),
					"eventID":   c.Get("eventID").(string),
				})
		}
		logger.Info("Succefuly deleted template", append(loggingFields, logging.String("id", templateId))...)

		return c.JSON(http.StatusOK, templateId)
	}
}
