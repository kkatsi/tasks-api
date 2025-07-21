package storage

import (
	"context"
	"rest-api/internal/storage/db"
	"rest-api/internal/utils"
	"sync"
	"time"
)

// In-memory implementation
type MemoryStore struct {
	mu    sync.RWMutex
	tasks map[string]db.Task
}

func NewMemoryStore() Storage {
	return &MemoryStore{
		tasks: make(map[string]db.Task),
	}
}

// Create implements Storage.
func (s *MemoryStore) Create(ctx context.Context, task *db.Task) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks[task.ID] = *task
	return task.ID, nil
}

// Delete implements Storage.
func (s *MemoryStore) Delete(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, taskExists := s.tasks[id]; !taskExists {
		return utils.ErrTaskNotFound
	}

	delete(s.tasks, id)
	return nil
}

// Get implements Storage.
func (s *MemoryStore) Get(ctx context.Context, id string) (*db.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, taskExists := s.tasks[id]
	if !taskExists {
		return nil, utils.ErrTaskNotFound
	}

	return &task, nil
}

// GetAll implements Storage.
func (s *MemoryStore) GetAll(ctx context.Context) ([]db.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]db.Task, 0, len(s.tasks))

	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// Update implements Storage.
func (s *MemoryStore) Update(ctx context.Context, id string, task *db.Task) (*db.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return nil, utils.ErrTaskNotFound
	}

	task.CreatedAt = s.tasks[id].CreatedAt
	task.UpdatedAt = time.Now()

	s.tasks[id] = *task

	updatedTask := s.tasks[id]

	return &updatedTask, nil
}
