package storage

import (
	"context"
	"database/sql"
	"fmt"
	"rest-api/internal/storage/db"
	"time"
)

type SQLiteStore struct {
	db      *sql.DB
	queries *db.Queries // sqlc generated
}

func NewSQLiteStore(database *sql.DB) Storage {
	return &SQLiteStore{
		db:      database,
		queries: db.New(database),
	}
}

// Create implements Storage.
func (s *SQLiteStore) Create(ctx context.Context, task *db.Task) (string, error) {
	createdTask, err := s.queries.CreateTask(ctx, db.CreateTaskParams{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	})

	if err != nil {
		return "", fmt.Errorf("internal server error")
	}

	return createdTask.ID, nil
}

// Delete implements Storage.
func (s *SQLiteStore) Delete(ctx context.Context, id string) error {
	_, err := s.queries.DeleteTask(ctx, id)

	if err != nil {
		return fmt.Errorf("internal server error")
	}

	return nil
}

// Get implements Storage.
func (s *SQLiteStore) Get(ctx context.Context, id string) (*db.Task, error) {
	task, err := s.queries.GetTask(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	return &task, nil
}

// GetAll implements Storage.
func (s *SQLiteStore) GetAll(ctx context.Context) ([]db.Task, error) {
	tasks, err := s.queries.GetTasks(ctx)

	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	return tasks, nil
}

// Update implements Storage.
func (s *SQLiteStore) Update(ctx context.Context, id string, task *db.Task) (*db.Task, error) {
	updatedTask, err := s.queries.UpdateTask(ctx, db.UpdateTaskParams{
		ID:          id,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		UpdatedAt:   time.Now(),
	})

	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	return &updatedTask, nil
}
