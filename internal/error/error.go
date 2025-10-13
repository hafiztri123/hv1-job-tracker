package appError

import (
	"errors"
	"net/http"
)

var (
	ErrDuplicateEmail = &AppError{
		Err:        errors.New("email already exists"),
		Message:    "Email already exists",
		StatusCode: http.StatusConflict,
	}

	ErrInvalidInput = &AppError{
		Err:        errors.New("invalid input"),
		Message:    "Invalid input provided",
		StatusCode: http.StatusBadRequest,
	}

	ErrNotFound = &AppError{
		Err:        errors.New("not found"),
		Message:    "Resource not found",
		StatusCode: http.StatusNotFound,
	}

	ErrUnauthorized = &AppError{
		Err:        errors.New("unauthorized"),
		Message:    "Unauthorized access",
		StatusCode: http.StatusUnauthorized,
	}
)

func NewNotFoundErr(errorMsg string) *AppError {
	return &AppError{
		Err:        errors.New(errorMsg),
		Message:    "Not Found",
		StatusCode: http.StatusNotFound,
	}
}

func NewInternalServerError(errorMsg string) *AppError {
	return &AppError{
		Err:        errors.New(errorMsg),
		Message:    "Internal Server Error",
		StatusCode: http.StatusInternalServerError,
	}
}

func NewBadRequestError(errorMsg string) *AppError {
	return &AppError{
		Err:        errors.New(errorMsg),
		Message:    "Bad Request",
		StatusCode: http.StatusBadRequest,
	}
}
