package service

import (
	"context"
	"rest-api/internal/model"
	"rest-api/internal/storage"
	"rest-api/internal/storage/db"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	store storage.Storage
}

func NewUserService(store storage.Storage) *UserService {
	return &UserService{
		store: store,
	}
}

func (s *UserService) Create(ctx context.Context, reqBody *model.RegisterUserRequest) (string, error) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(reqBody.Password), bcrypt.DefaultCost)

	newUser := db.User{
		ID:           uuid.NewString(),
		Username:     reqBody.Username,
		Email:        reqBody.Email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return s.store.CreateUser(ctx, &newUser)
}
