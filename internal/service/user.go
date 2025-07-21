package service

import "rest-api/internal/storage"

type UserService struct {
	store storage.Storage
}

func NewUserService(store storage.Storage) *UserService {
	return &UserService{
		store: store,
	}
}

func (s *UserService) Create() {
}
