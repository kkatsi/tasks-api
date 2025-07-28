package service

import (
	"context"
	"errors"
	"rest-api/internal/apperrors"
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

func (s *UserService) getUserId(ctx context.Context) (string, error) {
	userId, ok := ctx.Value("userId").(string)

	if !ok {
		return "", apperrors.ErrInternalError
	}

	if err := uuid.Validate(userId); err != nil {
		return "", apperrors.ErrInternalError
	}

	return userId, nil
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

func (s *UserService) GetMyUser(ctx context.Context) (*db.User, error) {
	userId, err := s.getUserId(ctx)

	if err != nil {
		return nil, err
	}

	return s.store.GetUserById(ctx, userId)
}

func (s *UserService) DeleteMyUser(ctx context.Context) error {
	userId, err := s.getUserId(ctx)

	if err != nil {
		return err
	}

	return s.store.DeleteUser(ctx, userId)
}

func (s *UserService) UpdateMyUser(ctx context.Context, reqBody *model.UpdateUserRequest) (*db.User, error) {
	userId, err := s.getUserId(ctx)

	if err != nil {
		return nil, err
	}

	if err := reqBody.Validate(); err != nil {
		return nil, err
	}

	existingUser, err := s.store.GetUserByEmail(ctx, reqBody.Email)

	if existingUser != nil && existingUser.ID != userId {
		return nil, apperrors.ErrUserEmailExists
	}

	if err != nil && !errors.Is(err, apperrors.ErrUserNotFound) {
		return nil, apperrors.ErrInternalError
	}

	existingUser, err = s.store.GetUserByUsername(ctx, reqBody.Username)

	if existingUser != nil && existingUser.ID != userId {
		return nil, apperrors.ErrUserUsernameExists
	}

	if err != nil && !errors.Is(err, apperrors.ErrUserNotFound) {
		return nil, apperrors.ErrInternalError
	}

	return s.store.UpdateUser(ctx, &db.User{
		ID:       userId,
		Username: reqBody.Username,
		Email:    reqBody.Email,
	})
}

func (s *UserService) UpdateMyPassword(ctx context.Context, reqBody *model.UpdatePasswordRequest) error {
	userId, err := s.getUserId(ctx)

	if err != nil {
		return err
	}

	if err := reqBody.Validate(); err != nil {
		return err
	}

	user, err := s.store.GetUserById(ctx, userId)

	if err != nil {
		return apperrors.ErrInternalError
	}

	existingPassword := user.PasswordHash

	err = bcrypt.CompareHashAndPassword([]byte(existingPassword), []byte(reqBody.OldPassword))

	if err != nil {
		return apperrors.ErrWrongPassword
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(reqBody.NewPassword), bcrypt.DefaultCost)

	return s.store.UpdatePassword(ctx, userId, string(hashedPassword))

}
