package errors

import (
	"net/http"

	"github.com/project_1/cmd/domain"
)

// NewBadRequestAppError create a new bad request error
func NewBadRequestAppError(message string) domain.AppError {
	return domain.AppError{message, "bad_request", http.StatusBadRequest, domain.CauseList{}}
}

// NewInternalServerAppError create a new internal server error
func NewInternalServerAppError(message string, err error) domain.AppError {
	cause := domain.CauseList{}
	if err != nil {
		cause = append(cause, err.Error())
	}
	return domain.AppError{message, "internal_server_error", http.StatusInternalServerError, cause}
}

// NewStatusNotFoundAppError create a new not found error
func NewStatusNotFoundAppError(message string) domain.AppError {
	return domain.AppError{message, "status_not_found", http.StatusNotFound, domain.CauseList{}}
}
