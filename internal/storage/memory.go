package storage

import (
	"context"
	"rest-api/internal/apperrors"
	"rest-api/internal/model"
	"rest-api/internal/storage/db"
	"sync"
	"time"
)

type MemoryStore struct {
	mu    sync.RWMutex
	tasks map[string]db.Task
}

func NewMemoryStore() Storage {
	return &MemoryStore{
		tasks: make(map[string]db.Task),
	}
}

func (s *MemoryStore) CreateTask(ctx context.Context, task *db.Task) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tasks[task.ID] = *task
	return task.ID, nil
}

func (s *MemoryStore) DeleteTask(ctx context.Context, id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, taskExists := s.tasks[id]; !taskExists {
		return apperrors.ErrTaskNotFound
	}

	delete(s.tasks, id)
	return nil
}

func (s *MemoryStore) GetTask(ctx context.Context, id string) (*db.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, taskExists := s.tasks[id]
	if !taskExists {
		return nil, apperrors.ErrTaskNotFound
	}

	return &task, nil
}

func (s *MemoryStore) GetAllTasks(ctx context.Context, paginatinonParams model.PaginationParams) ([]db.Task, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]db.Task, 0, len(s.tasks))

	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *MemoryStore) UpdateTask(ctx context.Context, id string, task *db.Task) (*db.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return nil, apperrors.ErrTaskNotFound
	}

	task.CreatedAt = s.tasks[id].CreatedAt
	task.UpdatedAt = time.Now()

	s.tasks[id] = *task

	updatedTask := s.tasks[id]

	return &updatedTask, nil
}

//user

func (s *MemoryStore) CreateUser(ctx context.Context, user *db.User) (string, error) {
	return "", nil
}

func (s *MemoryStore) GetUserByUsername(ctx context.Context, username string) (*db.User, error) {
	return nil, nil
}

func (s *MemoryStore) GetUserByEmail(ctx context.Context, username string) (*db.User, error) {
	return nil, nil
}

//auth

func (s *MemoryStore) GetRefreshToken(ctx context.Context, hashedRefreshToken string) (*db.RefreshToken, error) {
	return nil, nil
}

func (s *MemoryStore) CreateRefreshTokenRecord(ctx context.Context, refreshTokenEntity db.RefreshToken) error {
	return nil
}

func (s *MemoryStore) DeleteExpiredTokens(ctx context.Context) error {
	return nil
}

func (s *MemoryStore) DeleteRefreshToken(ctx context.Context, userId string, hashedRefreshToken string) error {
	return nil
}
