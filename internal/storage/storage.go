package storage

import (
	"context"
	"rest-api/internal/model"
	"rest-api/internal/storage/db"
)

type Storage interface {

	//task
	CreateTask(ctx context.Context, task *db.Task) (string, error)
	UpdateTask(ctx context.Context, id string, task *db.Task) (*db.Task, error)
	DeleteTask(ctx context.Context, id string) error
	GetTask(ctx context.Context, id string) (*db.Task, error)
	GetAllTasks(ctx context.Context, paginationParams model.PaginationParams) ([]db.Task, error)

	//user
	CreateUser(ctx context.Context, user *db.User) (string, error)
	GetUserByUsername(ctx context.Context, username string) (*db.User, error)
}
