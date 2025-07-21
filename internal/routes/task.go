package routes

import (
	"net/http"
	"rest-api/internal/handler"
)

func SetupTaskRoutes(mux *http.ServeMux, h *handler.TaskHandler) {
	mux.HandleFunc("GET /tasks", h.GetAll)
	mux.HandleFunc("GET /tasks/{id}", h.Get)
	mux.HandleFunc("POST /tasks", h.Create)
	mux.HandleFunc("DELETE /tasks/{id}", h.Delete)
	mux.HandleFunc("PUT /tasks/{id}", h.Update)
}
