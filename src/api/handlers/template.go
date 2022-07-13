package handlers

import (
	"net/http"

	"github.com/allanassis/reddere/src/api/errors"
	"github.com/allanassis/reddere/src/api/response"
	"github.com/allanassis/reddere/src/observability/logging"
	"github.com/allanassis/reddere/src/services"
	"github.com/allanassis/reddere/src/services/entities"
	"github.com/allanassis/reddere/src/storages"
	"github.com/labstack/echo/v4"
)

// ShowAccount godoc
// @Summary      Save a template
// @Description  post template
// @Tags         template
// @Accept       json
// @Produce      json
// @Param        template  body     entities.Template   true  "Template"
// @Success      200  {object}  response.ApiResponse
// @Failure      400  {object}  response.ApiResponse
// @Failure      404  {object}  response.ApiResponse
// @Failure      500  {object}  response.ApiResponse
// @Router       /template [post]
func PostTemplate(service services.Service, storage storages.Storage, logger *logging.Logger) func(c echo.Context) error {
	return func(c echo.Context) error {
		template := new(entities.Template)

		loggingFields := []logging.Field{logging.String("entity", "Template")}
		err := c.Bind(template)
		if err != nil {

			logger.Error(string(errors.API_BIND_PAYLOAD_ERROR),
				append(loggingFields, logging.String("error", err.Error()))...,
			)

			resp := response.NewApiResponse(
				response.WithEventID(c.Get("eventID").(string)),
				response.WithError(errors.API_BIND_PAYLOAD_ERROR),
			)

			return c.JSON(http.StatusUnprocessableEntity, resp)
		}
		loggingFields = append(loggingFields, logging.Any("template", template))
		logger.Debug("Received request with payload", loggingFields...)

		isValid, err := template.IsValid()
		if !isValid {

			logger.Error(string(errors.API_BIND_PAYLOAD_ERROR),
				append(loggingFields, logging.String("error", err.Error()))...,
			)

			resp := response.NewApiResponse(
				response.WithEventID(c.Get("eventID").(string)),
				response.WithError(errors.API_BIND_PAYLOAD_ERROR),
				response.WithData(err),
			)

			return c.JSON(http.StatusBadRequest, resp)
		}
		logger.Debug("Fields succefuly validated", loggingFields...)

		id, err := service.Save(template)
		if err != nil {
			logger.Error(string(errors.API_POST_ERROR),
				append(loggingFields, logging.String("error", err.Error()))...,
			)

			resp := response.NewApiResponse(
				response.WithEventID(c.Get("eventID").(string)),
				response.WithError(errors.API_POST_ERROR),
			)

			return c.JSON(http.StatusInternalServerError, resp)
		}
		logger.Info("Succefuly saved template", append(loggingFields, logging.String("id", id))...)

		template.ID = id

		resp := response.NewApiResponse(
			response.WithEventID(c.Get("eventID").(string)),
			response.WithMessage("Succefuly saved template"),
			response.WithData(template),
		)
		return c.JSON(http.StatusCreated, resp)
	}
}

// ShowTemplate godoc
// @Summary      Show a template
// @Description  get template by ID
// @Tags         template
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Template ID"
// @Success      200  {object}  response.ApiResponse
// @Failure      400  {object}  response.ApiResponse
// @Failure      404  {object}  response.ApiResponse
// @Failure      500  {object}  response.ApiResponse
// @Router       /template/{id} [get]
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
			resp := response.NewApiResponse(
				response.WithEventID(c.Get("eventID").(string)),
				response.WithError(errors.API_GET_ERROR),
			)

			return c.JSON(http.StatusInternalServerError, resp)
		}
		loggingFields = append(loggingFields, logging.Any("template", template))
		logger.Info("Succefuly get template", append(loggingFields, logging.String("id", templateId))...)

		resp := response.NewApiResponse(
			response.WithEventID(c.Get("eventID").(string)),
			response.WithMessage("Succefuly get template"),
			response.WithData(template),
		)

		return c.JSON(http.StatusOK, resp)
	}
}

// DeleteTemplate godoc
// @Summary      Delete a template
// @Description  delete template by ID
// @Tags         template
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "Template ID"
// @Success      200  {object}  response.ApiResponse
// @Failure      400  {object}  response.ApiResponse
// @Failure      404  {object}  response.ApiResponse
// @Failure      500  {object}  response.ApiResponse
// @Router       /template/{id} [delete]
func DeleteTemplate(service services.Service, storage storages.Storage, logger *logging.Logger) func(c echo.Context) error {
	return func(c echo.Context) error {
		loggingFields := []logging.Field{logging.String("entity", "Template")}

		templateId := c.Param("id")

		err := service.Delete(templateId, "template")
		if err != nil {
			logger.Error(string(errors.API_DELETE_ERROR),
				append(loggingFields, logging.String("error", err.Error()))...,
			)

			resp := response.NewApiResponse(
				response.WithEventID(c.Get("eventID").(string)),
				response.WithError(errors.API_DELETE_ERROR),
			)

			return c.JSON(http.StatusInternalServerError, resp)
		}
		logger.Info("Succefuly deleted template", append(loggingFields, logging.String("id", templateId))...)

		template := &entities.Template{
			ID: templateId,
		}

		resp := response.NewApiResponse(
			response.WithEventID(c.Get("eventID").(string)),
			response.WithMessage("Succefuly deleted template"),
			response.WithData(template),
		)

		return c.JSON(http.StatusOK, resp)
	}
}
