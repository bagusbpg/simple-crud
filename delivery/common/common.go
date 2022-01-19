package common

import (
	"net/http"
)

type simpleResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SimpleResponse(code int, message string, data interface{}) simpleResponse {
	switch code {
	case http.StatusForbidden:
		if message == "" {
			message = "status forbidden"
		}
	case http.StatusInternalServerError:
		if message == "" {
			message = "status internal server error"
		}
	case http.StatusBadRequest:
		if message == "" {
			message = "status bad response"
		}
	case http.StatusUnauthorized:
		if message == "" {
			message = "status unauthorized"
		}
	default:
		if message == "" {
			message = "status ok"
		}
	}
	return simpleResponse{code, message, data}
}
