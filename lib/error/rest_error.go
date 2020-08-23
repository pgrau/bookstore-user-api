package error

import (
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status int `json:"status"`
	Error string `json:"error"`
}

func BadRequest(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad_request",
	}
}

func NotFound(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error: "not_found",
	}
}

func Conflict(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusConflict,
		Error: "conflict",
	}
}

func InternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusInternalServerError,
		Error: "internal_server_error",
	}
}