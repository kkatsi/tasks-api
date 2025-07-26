package handler

import (
	"errors"
	"net/http"
	"rest-api/internal/apperrors"
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
	case errors.Is(err, apperrors.ErrInvalidID), errors.Is(err, apperrors.ErrInvalidRequestBody):
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, apperrors.ErrUserNotFound):
		utils.ErrorResponse(w, http.StatusNotFound, err.Error())
	default:
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *UserHandler) GetMyUser(w http.ResponseWriter, r *http.Request)        {}
func (h *UserHandler) DeleteMyUser(w http.ResponseWriter, r *http.Request)     {}
func (h *UserHandler) UpdateMyUser(w http.ResponseWriter, r *http.Request)     {}
func (h *UserHandler) UpdateMyPassword(w http.ResponseWriter, r *http.Request) {}
