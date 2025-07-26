package apperrors

import "errors"

var (
	ErrTaskNotFound        = errors.New("task not found")
	ErrTitleRequired       = errors.New("title is required")
	ErrDescriptionRequired = errors.New("description is required")
	ErrCompletedRequired   = errors.New("completed field is required")
)
