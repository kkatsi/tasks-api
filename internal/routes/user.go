package routes

import (
	"net/http"
	"rest-api/internal/handler"
)

func SetupUserRoutes(mux *http.ServeMux, h *handler.UserHandler) {
	mux.HandleFunc("GET /users/me", h.GetMyUser)
	mux.HandleFunc("DELETE /users/me", h.DeleteMyUser)
	mux.HandleFunc("PUT /users/me", h.UpdateMyUser)
	mux.HandleFunc("PUT /users/me/password", h.UpdateMyPassword)
}
