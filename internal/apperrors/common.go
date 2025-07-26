package apperrors

import "errors"

var (
	ErrInvalidID          = errors.New("invalid task ID")
	ErrInvalidRequestBody = errors.New("invalid request body")
	ErrInvalidEmail       = errors.New("invalid email")
)
