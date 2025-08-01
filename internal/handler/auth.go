package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"rest-api/internal/apperrors"
	"rest-api/internal/model"
	"rest-api/internal/service"
	"rest-api/internal/utils"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func handleAuthError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, apperrors.ErrInvalidID), errors.Is(err, apperrors.ErrInvalidCredentials), errors.Is(err, apperrors.ErrInvalidRequestBody), errors.Is(err, apperrors.ErrUsernameRequired), errors.Is(err, apperrors.ErrPasswordMinLength), errors.Is(err, apperrors.ErrEmailRequired), errors.Is(err, apperrors.ErrUserUsernameExists), errors.Is(err, apperrors.ErrUserEmailExists), errors.Is(err, apperrors.ErrTokenInvalid), errors.Is(err, apperrors.ErrTokenExpired):
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, apperrors.ErrUserNotFound):
		utils.ErrorResponse(w, http.StatusNotFound, err.Error())
	default:
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var reqBody model.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		handleAuthError(w, err)
		return
	}

	tokens, err := h.service.Register(ctx, &reqBody)

	if err != nil {
		handleAuthError(w, err)
		return
	}
	utils.JsonResponse(w, http.StatusCreated, tokens)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var reqBody model.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		handleAuthError(w, err)
		return
	}

	userWithTokens, err := h.service.Login(ctx, &reqBody)

	if err != nil {
		handleAuthError(w, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, userWithTokens)
}

func (h *AuthHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {}
func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request)  {}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var reqBody model.RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		handleAuthError(w, err)
		return
	}

	freshTokens, err := h.service.Refresh(ctx, &reqBody)

	if err != nil {
		handleAuthError(w, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, freshTokens)

}
