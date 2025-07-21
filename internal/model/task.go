package model

import (
	"rest-api/internal/utils"
	"time"
)

type Task struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   *bool  `json:"completed"`
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
