package user

import "errors"

var (
	ErrDuplicateEmail = errors.New("email already exists")
	ErrDatabase       = errors.New("database error")
)
