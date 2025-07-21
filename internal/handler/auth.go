package handler

import (
	"net/http"
	"rest-api/internal/service"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request)       {}
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request)          {}
func (h *AuthHandler) ForgotPassword(w http.ResponseWriter, r *http.Request) {}
func (h *AuthHandler) ResetPassword(w http.ResponseWriter, r *http.Request)  {}
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request)   {}
