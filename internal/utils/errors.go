package utils

import (
	"errors"
)

var (
	ErrTaskNotFound          = errors.New("task not found")
	ErrInvalidID             = errors.New("invalid task ID")
	ErrInvalidRequestBody    = errors.New("invalid request body")
	ErrTitleIsRequired       = errors.New("title is required")
	ErrDescriptionIsRequired = errors.New("description is required")
	ErrCompletedIsRequired   = errors.New("completed field is required")
)
