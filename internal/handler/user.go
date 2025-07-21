package handler

import (
	"net/http"
	"rest-api/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetMyUser(w http.ResponseWriter, r *http.Request)        {}
func (h *UserHandler) DeleteMyUser(w http.ResponseWriter, r *http.Request)     {}
func (h *UserHandler) UpdateMyUser(w http.ResponseWriter, r *http.Request)     {}
func (h *UserHandler) UpdateMyPassword(w http.ResponseWriter, r *http.Request) {}
