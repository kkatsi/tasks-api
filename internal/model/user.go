package model

import "rest-api/internal/apperrors"

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UpdateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UpdatePasswordRequest struct {
	NewPassword string `json:"new"`
	OldPassword string `json:"old"`
}

func (r *UpdateUserRequest) Validate() error {
	if r.Username == "" {
		return apperrors.ErrUsernameRequired
	}
	if r.Email == "" {
		return apperrors.ErrEmailRequired
	}
	if !IsValidEmail(r.Email) {
		return apperrors.ErrInvalidEmail
	}
	return nil
}

func (r *UpdatePasswordRequest) Validate() error {
	if len(r.NewPassword) < 8 {
		return apperrors.ErrPasswordMinLength
	}

	if r.NewPassword == r.OldPassword {
		return apperrors.ErrPasswordMatch
	}

	return nil
}
