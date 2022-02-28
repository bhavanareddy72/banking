package errs

import "net/http"

type AppError struct {
	code    int    `json:",omitempty"`
	message string `json:",message"`
}

func (e AppError) Asmessage() *AppError {
	return &AppError{
		message: e.message,
	}
}

func NewNotFounderror(message string) *AppError {
	return &AppError{
		message: message,
		code:    http.StatusNotFound,
	}
}
func NewValidationerror(message string) *AppError {
	return &AppError{
		message: message,
		code:    http.StatusUnprocessableEntity,
	}
}
