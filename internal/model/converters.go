package model

import (
	"rest-api/internal/storage/db"
)

func TaskToDTO(task db.Task) GetTaskDTO {
	return GetTaskDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   &task.Completed,
	}
}

func TasksToListDTO(tasks []db.Task) ListDTO[GetTaskDTO] {
	tasksResponse := make([]GetTaskDTO, 0, len(tasks))
	for _, task := range tasks {
		tasksResponse = append(tasksResponse, TaskToDTO(task))
	}

	return ListDTO[GetTaskDTO]{
		Data:  tasksResponse,
		Total: len(tasks),
		Pages: 1,
	}
}
