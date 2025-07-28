package model

import "rest-api/internal/apperrors"

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   *bool  `json:"completed"`
}

type TaskResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   *bool  `json:"completed"`
}

type TasksResponse[T any] struct {
	Data  []T `json:"data"`
	Total int `json:"total"`
	Pages int `json:"pages"`
}

func (r *CreateTaskRequest) Validate() error {
	if r.Title == "" {
		return apperrors.ErrTitleRequired
	}

	if r.Description == "" {
		return apperrors.ErrDescriptionRequired
	}
	return nil
}

func (r *UpdateTaskRequest) Validate() error {
	if r.Title == "" {
		return apperrors.ErrTitleRequired
	}

	if r.Description == "" {
		return apperrors.ErrDescriptionRequired
	}

	if r.Completed == nil {
		return apperrors.ErrCompletedRequired
	}
	return nil
}
