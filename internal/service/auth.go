package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"rest-api/internal/apperrors"
	"rest-api/internal/config"
	"rest-api/internal/model"
	"rest-api/internal/storage"
	"rest-api/internal/storage/db"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func generateAccessToken(userId string) (string, error) {

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(10 * time.Minute).Unix(),
		"iat":    time.Now().Unix(),
	})

	accessTokenString, err := accessToken.SignedString([]byte(config.Get().AccessSecret))

	if err != nil {
		return "", apperrors.ErrInternalError
	}

	return accessTokenString, nil

}

func (s *AuthService) generateRefreshToken(ctx context.Context, userId string) (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)

	if err != nil {
		return "", apperrors.ErrInternalError
	}

	refreshToken := base64.URLEncoding.EncodeToString(b)
	hash := sha256.Sum256([]byte(refreshToken))
	hashedToken := hex.EncodeToString(hash[:])

	err = s.storage.CreateRefreshTokenRecord(ctx, db.RefreshToken{
		ID:        uuid.NewString(),
		UserID:    userId,
		TokenHash: hashedToken,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	})

	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (s *AuthService) cleanupExpiredTokens(ctx context.Context) error {
	return s.storage.DeleteExpiredTokens(ctx)
}

func (s *AuthService) hashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

func (s *AuthService) Register(ctx context.Context, reqBody *model.RegisterUserRequest) (*model.TokensResponse, error) {
	if err := reqBody.Validate(); err != nil {
		return nil, err
	}

	createdUserID, err := s.userService.Create(ctx, reqBody)

	if err != nil {
		return nil, err
	}

	accessToken, err := generateAccessToken(createdUserID)

	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateRefreshToken(ctx, createdUserID)

	if err != nil {
		return nil, err
	}

	return &model.TokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}

func (s *AuthService) Login(ctx context.Context, reqBody *model.LoginUserRequest) (*model.LoginResponse, error) {
	err := s.cleanupExpiredTokens(ctx)

	if err != nil {
		return nil, apperrors.ErrInternalError
	}

	user, err := s.storage.GetUserByEmail(ctx, reqBody.Email)

	if user == nil {
		return nil, apperrors.ErrInvalidCredentials
	}

	if err != nil {
		return nil, apperrors.ErrInternalError
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(reqBody.Password))

	if err != nil {
		return nil, apperrors.ErrInvalidCredentials
	}

	accessToken, err := generateAccessToken(user.ID)

	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateRefreshToken(ctx, user.ID)

	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		User: model.UserResponse{
			Username: user.Username,
			Email:    user.Email,
		},
		Tokens: model.TokensResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil

}

func (s *AuthService) Refresh(ctx context.Context, reqBody *model.RefreshRequest) (*model.TokensResponse, error) {
	err := s.cleanupExpiredTokens(ctx)

	if err != nil {
		return nil, apperrors.ErrInternalError
	}

	providedHashedRefreshToken := s.hashToken(reqBody.RefreshToken)

	refreshTokenEntity, err := s.storage.GetRefreshToken(ctx, providedHashedRefreshToken)

	if err != nil {
		return nil, err
	}

	userId := refreshTokenEntity.UserID

	err = s.storage.DeleteRefreshToken(ctx, userId, providedHashedRefreshToken)

	if err != nil {
		return nil, apperrors.ErrInternalError
	}

	accessToken, err := generateAccessToken(userId)

	if err != nil {
		return nil, apperrors.ErrInternalError
	}

	refreshToken, err := s.generateRefreshToken(ctx, userId)

	if err != nil {
		return nil, apperrors.ErrInternalError
	}

	return &model.TokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
