package routes

import (
	"net/http"
	"rest-api/internal/handler"
)

func SetupTaskRoutes(mux *http.ServeMux, h *handler.TaskHandler, auth func(http.HandlerFunc) http.HandlerFunc) {
	handle := func(pattern string, handler http.HandlerFunc) {
		mux.HandleFunc(pattern, auth(handler))
	}

	handle("GET /tasks", h.GetAll)
	handle("GET /tasks/{id}", h.Get)
	handle("POST /tasks", h.Create)
	handle("DELETE /tasks/{id}", h.Delete)
	handle("PUT /tasks/{id}", h.Update)
}
