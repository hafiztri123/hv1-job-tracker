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
