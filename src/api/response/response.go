package response

import (
	"github.com/allanassis/reddere/src/api/errors"
)

var (
	SUCCESS = "success"
	ERROR   = "error"
)

type ApiResponse struct {
	Status    string      `json:"status" binding:"required"`
	ErrorCode string      `json:"errorCode,omitempty"`
	Message   string      `json:"message" binding:"required"`
	Data      interface{} `json:"data,omitempty"`
	EventID   string      `json:"eventId" binding:"required"`
}

type Option func(response *ApiResponse)

func WithError(err errors.ApiError) func(response *ApiResponse) {
	return func(response *ApiResponse) {
		response.Status = ERROR

		response.ErrorCode = err.String()
		response.Message = string(err)
	}
}

func WithMessage(message string) func(response *ApiResponse) {
	return func(response *ApiResponse) {
		response.Message = message
	}
}

func WithEventID(eventID string) func(response *ApiResponse) {
	return func(response *ApiResponse) {
		response.EventID = eventID
	}
}

func WithData(data interface{}) func(response *ApiResponse) {
	return func(response *ApiResponse) {
		response.Data = data
	}
}

func NewApiResponse(options ...Option) *ApiResponse {
	response := &ApiResponse{
		Status: SUCCESS,
	}

	for _, opt := range options {
		opt(response)
	}
	return response
}
