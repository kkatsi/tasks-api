package service

import (
	"context"
	"rest-api/internal/model"
	"rest-api/internal/storage"
	"rest-api/internal/storage/db"
	"rest-api/internal/utils"
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
	return s.store.Create(ctx, &task)
}

func (s *TaskService) Get(ctx context.Context, id string) (*db.Task, error) {
	if err := uuid.Validate(id); err != nil {
		return nil, utils.ErrInvalidID
	}
	return s.store.Get(ctx, id)
}

func (s *TaskService) GetAll(ctx context.Context, paginationParams model.PaginationParams) ([]db.Task, error) {

	return s.store.GetAll(ctx, paginationParams)
}

func (s *TaskService) Delete(ctx context.Context, id string) error {
	if err := uuid.Validate(id); err != nil {
		return utils.ErrInvalidID
	}

	return s.store.Delete(ctx, id)
}

func (s *TaskService) Update(ctx context.Context, id string, reqBody *model.UpdateTaskRequest) (*db.Task, error) {
	if err := uuid.Validate(id); err != nil {
		return nil, utils.ErrInvalidID
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

	return s.store.Update(ctx, id, &task)
}
