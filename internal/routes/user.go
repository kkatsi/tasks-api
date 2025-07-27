package routes

import (
	"net/http"
	"rest-api/internal/handler"
)

func SetupUserRoutes(mux *http.ServeMux, h *handler.UserHandler, auth func(http.HandlerFunc) http.HandlerFunc) {
	handle := func(pattern string, handler http.HandlerFunc) {
		mux.HandleFunc(pattern, auth(handler))
	}

	handle("GET /users/me", h.GetMyUser)
	handle("DELETE /users/me", h.DeleteMyUser)
	handle("PUT /users/me", h.UpdateMyUser)
	handle("PUT /users/me/password", h.UpdateMyPassword)
}
