package apperrors

import "errors"

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrUserEmailExists    = errors.New("email already exists")
	ErrUserUsernameExists = errors.New("username already exists")
	ErrUsernameRequired   = errors.New("username is required")
	ErrEmailRequired      = errors.New("email is required")
	ErrPasswordMatch      = errors.New("new and old passwords should not match")
	ErrWrongPassword      = errors.New("password is wrong")
	ErrPasswordMinLength  = errors.New("password minimum length is 8 characters")
)
