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

	userId, ok := ctx.Value("userId").(string)

	if !ok {
		return "", apperrors.ErrInternalError
	}

	task := db.Task{
		ID:          uuid.NewString(),
		Title:       reqBody.Title,
		Description: reqBody.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		UserID:      userId,
	}
	return s.store.CreateTask(ctx, &task)
}

func (s *TaskService) Get(ctx context.Context, id string) (*db.Task, error) {
	if err := uuid.Validate(id); err != nil {
		return nil, apperrors.ErrInvalidID
	}
	userId, ok := ctx.Value("userId").(string)

	if !ok {
		return nil, apperrors.ErrInternalError
	}
	return s.store.GetTask(ctx, id, userId)
}

func (s *TaskService) GetAll(ctx context.Context, paginationParams model.PaginationParams) ([]db.Task, error) {

	if err := paginationParams.Validate(); err != nil {
		return nil, err
	}

	userId, ok := ctx.Value("userId").(string)

	if !ok {
		return nil, apperrors.ErrInternalError
	}

	return s.store.GetAllTasks(ctx, userId, paginationParams)
}

func (s *TaskService) Delete(ctx context.Context, id string) error {
	if err := uuid.Validate(id); err != nil {
		return apperrors.ErrInvalidID
	}

	userId, ok := ctx.Value("userId").(string)

	if !ok {
		return apperrors.ErrInternalError
	}

	return s.store.DeleteTask(ctx, id, userId)
}

func (s *TaskService) Update(ctx context.Context, id string, reqBody *model.UpdateTaskRequest) (*db.Task, error) {
	if err := uuid.Validate(id); err != nil {
		return nil, apperrors.ErrInvalidID
	}

	if err := reqBody.Validate(); err != nil {
		return nil, err
	}

	userId, ok := ctx.Value("userId").(string)

	if !ok {
		return nil, apperrors.ErrInternalError
	}

	task := db.Task{
		ID:          id,
		Title:       reqBody.Title,
		Description: reqBody.Description,
		Completed:   *reqBody.Completed,
		UpdatedAt:   time.Now(),
		UserID:      userId,
	}

	return s.store.UpdateTask(ctx, &task)
}
