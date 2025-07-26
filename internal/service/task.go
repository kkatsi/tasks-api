package service

import (
	"context"
	"rest-api/internal/apperrors"
	"rest-api/internal/model"
	"rest-api/internal/storage"
	"rest-api/internal/storage/db"
	"time"

	"github.com/google/uuid"
)

// TaskService handles HTTP requests for tasks
type TaskService struct {
	store storage.Storage // Dependency injection
}

// Constructor
func NewTaskService(store storage.Storage) *TaskService {
	return &TaskService{
		store: store,
	}
}

func (s *TaskService) Create(ctx context.Context, reqBody model.CreateTaskRequest) (string, error) {

	if err := reqBody.Validate(); err != nil {
		return "", err
	}

	task := db.Task{
		ID:          uuid.NewString(),
		Title:       reqBody.Title,
		Description: reqBody.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return s.store.CreateTask(ctx, &task)
}

func (s *TaskService) Get(ctx context.Context, id string) (*db.Task, error) {
	if err := uuid.Validate(id); err != nil {
		return nil, apperrors.ErrInvalidID
	}
	return s.store.GetTask(ctx, id)
}

func (s *TaskService) GetAll(ctx context.Context, paginationParams model.PaginationParams) ([]db.Task, error) {

	if err := paginationParams.Validate(); err != nil {
		return nil, err
	}

	return s.store.GetAllTasks(ctx, paginationParams)
}

func (s *TaskService) Delete(ctx context.Context, id string) error {
	if err := uuid.Validate(id); err != nil {
		return apperrors.ErrInvalidID
	}

	return s.store.DeleteTask(ctx, id)
}

func (s *TaskService) Update(ctx context.Context, id string, reqBody *model.UpdateTaskRequest) (*db.Task, error) {
	if err := uuid.Validate(id); err != nil {
		return nil, apperrors.ErrInvalidID
	}

	if err := reqBody.Validate(); err != nil {
		return nil, err
	}

	task := db.Task{
		ID:          id,
		Title:       reqBody.Title,
		Description: reqBody.Description,
		Completed:   *reqBody.Completed,
		UpdatedAt:   time.Now(),
	}

	return s.store.UpdateTask(ctx, id, &task)
}
