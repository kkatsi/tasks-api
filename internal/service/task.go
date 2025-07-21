package service

import (
	"rest-api/internal/model"
	"rest-api/internal/storage"
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

func (s *TaskService) Create(reqBody model.CreateTaskRequest) (string, error) {

	if err := reqBody.Validate(); err != nil {
		return "", err
	}

	task := model.Task{
		Id:          uuid.NewString(),
		Title:       reqBody.Title,
		Description: reqBody.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return s.store.Create(&task)
}

func (s *TaskService) Get(id string) (*model.Task, error) {
	if err := uuid.Validate(id); err != nil {
		return nil, utils.ErrInvalidID
	}
	return s.store.Get(id)
}

func (s *TaskService) GetAll() ([]*model.Task, error) {
	return s.store.GetAll()
}

func (s *TaskService) Delete(id string) error {
	if err := uuid.Validate(id); err != nil {
		return utils.ErrInvalidID
	}

	return s.store.Delete(id)
}

func (s *TaskService) Update(id string, reqBody *model.UpdateTaskRequest) (*model.Task, error) {
	if err := uuid.Validate(id); err != nil {
		return nil, utils.ErrInvalidID
	}

	if err := reqBody.Validate(); err != nil {
		return nil, err
	}

	task := model.Task{
		Id:          id,
		Title:       reqBody.Title,
		Description: reqBody.Description,
		Completed:   *reqBody.Completed,
		UpdatedAt:   time.Now(),
	}

	return s.store.Update(id, &task)
}
