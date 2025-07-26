package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"rest-api/internal/apperrors"
	"rest-api/internal/model"
	"rest-api/internal/service"
	"rest-api/internal/utils"
	"strconv"
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

func handleTaskError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, apperrors.ErrInvalidID), errors.Is(err, apperrors.ErrInvalidRequestBody), errors.Is(err, apperrors.ErrTitleRequired), errors.Is(err, apperrors.ErrDescriptionRequired), errors.Is(err, apperrors.ErrCompletedRequired):
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, apperrors.ErrTaskNotFound):
		utils.ErrorResponse(w, http.StatusNotFound, err.Error())
	default:
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var reqBody model.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		handleTaskError(w, err)
		return
	}

	taskId, err := h.service.Create(ctx, reqBody)

	if err != nil {
		handleTaskError(w, err)
		return
	}
	utils.JsonResponse(w, http.StatusCreated, map[string]string{"id": taskId})

}

func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	taskId := r.PathValue("id")
	task, err := h.service.Get(ctx, taskId)

	if err != nil {
		handleTaskError(w, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, model.TaskToDTO(*task))
}

func (h *TaskHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil && limitStr != "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "limit must be a valid number")
		return

	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil && offsetStr != "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "offset must be a valid number")
		return
	}

	tasks, err := h.service.GetAll(ctx, model.PaginationParams{
		Limit:  limit,
		Offset: offset,
	})

	if err != nil {
		handleTaskError(w, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, model.TasksToListDTO(tasks))
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	taskId := r.PathValue("id")
	err := h.service.Delete(ctx, taskId)

	if err != nil {
		handleTaskError(w, err)
		return
	}
	utils.JsonResponse(w, http.StatusNoContent, nil)
}

func (h *TaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	taskId := r.PathValue("id")
	var reqBody model.UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, apperrors.ErrInvalidRequestBody.Error())
		return
	}

	task, err := h.service.Update(ctx, taskId, &reqBody)

	if err != nil {
		handleTaskError(w, err)
		return
	}

	utils.JsonResponse(w, http.StatusOK, model.TaskToDTO(*task))
}
