package resterrors

import (
	"encoding/json"
	"errors"
	"net/http"
)

// RestErr struct
type RestErr struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

// NewError factory
func NewError(msg string) error {
	return errors.New(msg)
}

// NewBadRequestError func
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NewNotFoundError func
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

// NewInternalServerError func
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

// NewRestErrorFromBytes func
func NewRestErrorFromBytes(bytes []byte) (*RestErr, error) {
	var apiErr RestErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return &apiErr, nil
}

// NewRestError return new RestErr
func NewRestError(message string, status int, err string, causes []interface{}) *RestErr {
	return &RestErr{
		Message: message,
		Status:  status,
		Error:   err,
		Causes:  causes,
	}
}
