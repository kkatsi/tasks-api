package storage

import (
	"context"
	"rest-api/internal/model"
	"rest-api/internal/storage/db"
)

type Storage interface {
	Create(ctx context.Context, task *db.Task) (string, error)
	Update(ctx context.Context, id string, task *db.Task) (*db.Task, error)
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*db.Task, error)
	GetAll(ctx context.Context, paginationParams model.PaginationParams) ([]db.Task, error)
}
