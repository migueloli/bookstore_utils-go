package resterrors

import (
	"net/http"
)

// RestErr is the base structure for error responses
type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

// NewBadRequestError is used to create a RestErr informing a BadRequest and a message.
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NewNotFoundError is used to create a RestErr informing a NotFound and a message.
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

// NewInternalServerError is used to create a RestErr informing a InternalServerError and a message.
func NewInternalServerError(message string, err error) *RestErr {
	result := &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}

	if err != nil {
		result.Causes = append(result.Causes, err.Error())
	}

	return result
}
