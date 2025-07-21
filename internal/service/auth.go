package service

import "rest-api/internal/storage"

type AuthService struct {
	storage storage.Storage
}

func NewAuthService(storage storage.Storage) *AuthService {
	return &AuthService{
		storage: storage,
	}
}
