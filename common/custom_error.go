package common

import "net/http"

type CustomError struct {
	StatusCode int    `json:"status,omitempty"`
	Message    string `json:"message"`
}

func (e CustomError) AsMessage() *CustomError {
	return &CustomError{
		Message: e.Message,
	}
}

func UnexpectedServerError(message string) *CustomError {
	return &CustomError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}

func NotFoundError(message string) *CustomError {
	return &CustomError{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

func BadRequestError(message string) *CustomError {
	return &CustomError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

func ConflictError(message string) *CustomError {
	return &CustomError{
		StatusCode: http.StatusConflict,
		Message:    message,
	}
}
