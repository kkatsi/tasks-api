package service

import (
	"context"
	"rest-api/internal/model"
	"rest-api/internal/storage"
)

type AuthService struct {
	userService *UserService
	storage     storage.Storage
}

func NewAuthService(storage storage.Storage, userService *UserService) *AuthService {
	return &AuthService{
		userService: userService,
		storage:     storage,
	}
}

func (s *AuthService) Register(ctx context.Context, reqBody *model.RegisterUserRequest) (string, error) {
	if err := reqBody.Validate(); err != nil {
		return "", err
	}

	createdUserID, err := s.userService.Create(ctx, reqBody)

	if err != nil {
		return "", err
	}

	//generate jwt

	return createdUserID, nil

}
