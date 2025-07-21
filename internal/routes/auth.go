package routes

import (
	"net/http"
	"rest-api/internal/handler"
)

func SetupAuthRoutes(mux *http.ServeMux, h *handler.AuthHandler) {
	mux.HandleFunc("POST /auth/register", h.Register)
	mux.HandleFunc("POST /auth/login", h.Login)
	mux.HandleFunc("POST /auth/forgot-password", h.ForgotPassword)
	mux.HandleFunc("POST /auth/reset-password", h.ResetPassword)
	mux.HandleFunc("POST /auth/refresh", h.RefreshToken)
}
