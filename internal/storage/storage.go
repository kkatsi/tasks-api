package storage

import (
	"context"
	"rest-api/internal/model"
	"rest-api/internal/storage/db"
)

type Storage interface {

	//task
	CreateTask(ctx context.Context, task *db.Task) (string, error)
	UpdateTask(ctx context.Context, task *db.Task) (*db.Task, error)
	DeleteTask(ctx context.Context, taskId string, userId string) error
	GetTask(ctx context.Context, taskId string, userId string) (*db.Task, error)
	GetAllTasks(ctx context.Context, userId string, paginationParams model.PaginationParams) ([]db.Task, error)

	//user
	CreateUser(ctx context.Context, user *db.User) (string, error)
	GetUserByUsername(ctx context.Context, username string) (*db.User, error)
	GetUserByEmail(ctx context.Context, email string) (*db.User, error)

	//auth
	GetRefreshToken(ctx context.Context, hashedRefreshToken string) (*db.RefreshToken, error)
	CreateRefreshTokenRecord(ctx context.Context, refreshTokenEntity db.RefreshToken) error
	DeleteExpiredTokens(ctx context.Context) error
	DeleteRefreshToken(ctx context.Context, userId string, hashedRefreshToken string) error
}
