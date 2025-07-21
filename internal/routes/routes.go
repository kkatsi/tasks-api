package routes

import (
	"net/http"
	"rest-api/internal/handler"
	"rest-api/internal/utils"
)

func SetupRoutes(mux *http.ServeMux, h *handler.TaskHandler) {

	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		utils.JsonResponse(w, http.StatusOK, map[string]string{"status": "OK"})
	})

	mux.HandleFunc("GET /tasks", h.GetAll)
	mux.HandleFunc("GET /tasks/{id}", h.Get)
	mux.HandleFunc("POST /tasks", h.Create)
	mux.HandleFunc("DELETE /tasks/{id}", h.Delete)
	mux.HandleFunc("PUT /tasks/{id}", h.Update)

}
