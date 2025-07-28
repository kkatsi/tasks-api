package model

import (
	"rest-api/internal/storage/db"
)

func TaskToTaskResponse(task db.Task) TaskResponse {
	return TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   &task.Completed,
	}
}

func TasksToList(tasks []db.Task) TasksResponse[TaskResponse] {
	tasksResponse := make([]TaskResponse, 0, len(tasks))
	for _, task := range tasks {
		tasksResponse = append(tasksResponse, TaskToTaskResponse(task))
	}

	return TasksResponse[TaskResponse]{
		Data:  tasksResponse,
		Total: len(tasks),
		Pages: 1,
	}
}
