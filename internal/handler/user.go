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

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func handleUserError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, apperrors.ErrInvalidID), errors.Is(err, apperrors.ErrInvalidRequestBody), errors.Is(err, apperrors.ErrPasswordMatch), errors.Is(err, apperrors.ErrWrongPassword), errors.Is(err, apperrors.ErrUserEmailExists), errors.Is(err, apperrors.ErrUserUsernameExists), errors.Is(err, apperrors.ErrUsernameRequired), errors.Is(err, apperrors.ErrPasswordMinLength), errors.Is(err, apperrors.ErrEmailRequired):
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, apperrors.ErrUserNotFound):
		utils.ErrorResponse(w, http.StatusNotFound, err.Error())
	default:
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *UserHandler) GetMyUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, err := h.service.GetMyUser(ctx)

	if err != nil {
		handleUserError(w, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, model.UserResponse{
		Username: user.Username,
		Email:    user.Email,
		ID:       user.ID,
	})
}
func (h *UserHandler) DeleteMyUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := h.service.DeleteMyUser(ctx)

	if err != nil {
		handleUserError(w, err)
		return
	}

	utils.JsonResponse(w, http.StatusNoContent, nil)
}
func (h *UserHandler) UpdateMyUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var reqBody model.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody.Error())
		return
	}

	updatedUser, err := h.service.UpdateMyUser(ctx, &reqBody)

	if err != nil {
		handleUserError(w, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, model.UserResponse{
		ID:       updatedUser.ID,
		Username: updatedUser.Username,
		Email:    updatedUser.Email,
	})

}
func (h *UserHandler) UpdateMyPassword(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var reqBody model.UpdatePasswordRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody.Error())
		return
	}

	err := h.service.UpdateMyPassword(ctx, &reqBody)

	if err != nil {
		handleUserError(w, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, nil)
}
