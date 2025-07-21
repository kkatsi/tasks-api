package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"rest-api/internal/model"
	"rest-api/internal/service"
	"rest-api/internal/utils"
)

// TaskHandler handles HTTP requests for tasks
type TaskHandler struct {
	service *service.TaskService // Dependency injection
}

// Constructor
func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, utils.ErrInvalidID), errors.Is(err, utils.ErrInvalidRequestBody), errors.Is(err, utils.ErrTitleIsRequired), errors.Is(err, utils.ErrDescriptionIsRequired), errors.Is(err, utils.ErrCompletedIsRequired):
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, utils.ErrTaskNotFound):
		utils.ErrorResponse(w, http.StatusNotFound, err.Error())
	default:
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var reqBody model.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		handleError(w, err)
		return
	}

	taskId, err := h.service.Create(ctx, reqBody)

	if err != nil {
		handleError(w, err)
		return
	}
	utils.JsonResponse(w, http.StatusCreated, map[string]string{"id": taskId})

}

func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	taskId := r.PathValue("id")
	task, err := h.service.Get(ctx, taskId)

	if err != nil {
		handleError(w, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, model.TaskToDTO(*task))
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := h.service.GetAll(ctx)

	if err != nil {
		handleError(w, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, model.TasksToListDTO(tasks))
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	taskId := r.PathValue("id")
	err := h.service.Delete(ctx, taskId)

	if err != nil {
		handleError(w, err)
		return
	}
	utils.JsonResponse(w, http.StatusNoContent, nil)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	taskId := r.PathValue("id")
	var reqBody model.UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, utils.ErrInvalidRequestBody.Error())
		return
	}

	task, err := h.service.Update(ctx, taskId, &reqBody)

	if err != nil {
		handleError(w, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, task)
}
