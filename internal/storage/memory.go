package storage

import (
	"rest-api/internal/model"
	"rest-api/internal/utils"
	"sync"
	"time"
)

type Storage interface {
	Create(task *model.Task) (string, error)
	Update(id string, task *model.Task) (*model.Task, error)
	Delete(id string) error
	Get(id string) (*model.Task, error)
	GetAll() ([]*model.Task, error)
}

// In-memory implementation
type MemoryStore struct {
	mu    sync.RWMutex
	tasks map[string]*model.Task
}

func NewMemoryStore() Storage {
	return &MemoryStore{
		tasks: make(map[string]*model.Task),
	}
}

// Create implements Storage.
func (s *MemoryStore) Create(task *model.Task) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks[task.Id] = task
	return task.Id, nil
}

// Delete implements Storage.
func (s *MemoryStore) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, taskExists := s.tasks[id]; !taskExists {
		return utils.ErrTaskNotFound
	}

	delete(s.tasks, id)
	return nil
}

// Get implements Storage.
func (s *MemoryStore) Get(id string) (*model.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, taskExists := s.tasks[id]; !taskExists {
		return nil, utils.ErrTaskNotFound
	}

	return s.tasks[id], nil
}

// GetAll implements Storage.
func (s *MemoryStore) GetAll() ([]*model.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]*model.Task, 0, len(s.tasks))

	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// Update implements Storage.
func (s *MemoryStore) Update(id string, task *model.Task) (*model.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return nil, utils.ErrTaskNotFound
	}

	task.CreatedAt = s.tasks[id].CreatedAt
	task.UpdatedAt = time.Now()

	s.tasks[id] = task

	return s.tasks[id], nil
}
