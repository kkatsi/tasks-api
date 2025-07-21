package model

import (
	"rest-api/internal/utils"
)

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   *bool  `json:"completed"`
}

type GetTaskDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   *bool  `json:"completed"`
}

type ListDTO[T any] struct {
	Data  []T `json:"data"`
	Total int `json:"total"`
	Pages int `json:"pages"`
}

func (r *CreateTaskRequest) Validate() error {
	if r.Title == "" {
		return utils.ErrTitleIsRequired
	}

	if r.Description == "" {
		return utils.ErrDescriptionIsRequired
	}
	return nil
}

func (r *UpdateTaskRequest) Validate() error {
	if r.Title == "" {
		return utils.ErrTitleIsRequired
	}

	if r.Description == "" {
		return utils.ErrDescriptionIsRequired
	}

	if r.Completed == nil {
		return utils.ErrCompletedIsRequired
	}
	return nil
}
