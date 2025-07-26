package storage

import (
	"context"
	"database/sql"
	"errors"
	"rest-api/internal/apperrors"
	"rest-api/internal/model"
	"rest-api/internal/storage/db"
	"strings"
	"time"
)

type SQLiteStore struct {
	db      *sql.DB
	queries *db.Queries
}

func NewSQLiteStore(database *sql.DB) Storage {
	return &SQLiteStore{
		db:      database,
		queries: db.New(database),
	}
}

func (s *SQLiteStore) CreateTask(ctx context.Context, task *db.Task) (string, error) {
	createdTask, err := s.queries.CreateTask(ctx, db.CreateTaskParams{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	})

	if err != nil {
		return "", errors.New("internal server error")
	}

	return createdTask.ID, nil
}

func (s *SQLiteStore) DeleteTask(ctx context.Context, id string) error {
	_, err := s.queries.DeleteTask(ctx, id)

	if err != nil {
		var error error = errors.New("internal server error")
		if err == sql.ErrNoRows {
			error = apperrors.ErrTaskNotFound
		}
		return error
	}

	return nil
}

func (s *SQLiteStore) GetTask(ctx context.Context, id string) (*db.Task, error) {
	task, err := s.queries.GetTask(ctx, id)

	if err != nil {
		var error error = err
		if err == sql.ErrNoRows {
			error = apperrors.ErrTaskNotFound
		}

		return nil, error
	}

	return &task, nil
}

func (s *SQLiteStore) GetAllTasks(ctx context.Context, paginationParams model.PaginationParams) ([]db.Task, error) {
	tasks, err := s.queries.GetTasks(ctx, db.GetTasksParams{
		Limit:  int64(paginationParams.Limit),
		Offset: int64(paginationParams.Offset),
	})

	if err != nil {
		return nil, errors.New("internal server error")
	}

	return tasks, nil
}

func (s *SQLiteStore) UpdateTask(ctx context.Context, id string, task *db.Task) (*db.Task, error) {
	updatedTask, err := s.queries.UpdateTask(ctx, db.UpdateTaskParams{
		ID:          id,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		var error error = err
		if err == sql.ErrNoRows {
			error = apperrors.ErrTaskNotFound
		}

		return nil, error
	}

	return &updatedTask, nil
}

//user

func (s *SQLiteStore) CreateUser(ctx context.Context, user *db.User) (string, error) {

	createdUser, err := s.queries.CreateUser(ctx, db.CreateUserParams{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	})

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			if strings.Contains(err.Error(), "users.username") {
				return "", apperrors.ErrUserUsernameExists
			}
			if strings.Contains(err.Error(), "users.email") {
				return "", apperrors.ErrUserEmailExists
			}
		}
		return "", errors.New("internal server error")
	}

	return createdUser.ID, nil
}

func (s *SQLiteStore) GetUserByUsername(ctx context.Context, username string) (*db.User, error) {
	user, err := s.queries.GetUserByUsername(ctx, username)

	if err != nil {
		var error error = err
		if err == sql.ErrNoRows {
			error = apperrors.ErrUserNotFound
		}

		return nil, error
	}

	return &user, nil
}
