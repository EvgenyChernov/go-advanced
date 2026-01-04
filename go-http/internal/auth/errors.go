package auth

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrInvalidEmail      = errors.New("invalid email")
	ErrInvalidName       = errors.New("invalid name")
)
