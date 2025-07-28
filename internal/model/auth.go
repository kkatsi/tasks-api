package model

import (
	"rest-api/internal/apperrors"
)

type TokensResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginResponse struct {
	User   UserResponse   `json:"user"`
	Tokens TokensResponse `json:"tokens"`
}

type RegisterUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

func (r *RegisterUserRequest) Validate() error {
	if r.Username == "" {
		return apperrors.ErrUsernameRequired
	}
	if r.Email == "" {
		return apperrors.ErrEmailRequired
	}
	if !IsValidEmail(r.Email) {
		return apperrors.ErrInvalidEmail
	}
	if len(r.Password) < 8 {
		return apperrors.ErrPasswordMinLength
	}
	return nil
}
