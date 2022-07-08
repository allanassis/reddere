package handlers

import "fmt"

type ApiError string

func (e ApiError) String() string {
	switch e {
	case API_HEALTHCHECK_ERROR:
		return "API_HEALTHCHECK_ERROR"
	case API_BIND_PAYLOAD_ERROR:
		return "API_BIND_PAYLOAD_ERROR"
	case API_DELETE_ERROR:
		return "API_DELETE_ERROR"
	case API_GET_ERROR:
		return "API_GET_ERROR"
	case API_POST_ERROR:
		return "API_POST_ERROR"
	default:
		return fmt.Sprintf("%s", string(e))
	}
}

const (
	API_HEALTHCHECK_ERROR  ApiError = "Healthcheck Failed"
	API_BIND_PAYLOAD_ERROR ApiError = "Payload is invalid"
	API_DELETE_ERROR       ApiError = "Deleting data failed"
	API_GET_ERROR          ApiError = "Get data failed"
	API_POST_ERROR         ApiError = "Save data failed"
)
