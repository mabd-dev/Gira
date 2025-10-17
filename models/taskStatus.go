package models

import "fmt"

type TaskStatus string

var (
	TaskStatusTodo       TaskStatus = "Todo"
	TaskStatusInProgress TaskStatus = "In Progress"
	TaskStatusInReview   TaskStatus = "In Review"
	TaskStatusStaging    TaskStatus = "Staging"
	TaskStatusDone       TaskStatus = "Done"
)

var TasksInOrder = []TaskStatus{
	TaskStatusTodo, TaskStatusInProgress, TaskStatusInReview, TaskStatusDone,
}

func getTaskStatusFrom(s string) (TaskStatus, error) {
	switch s {
	case "To Do":
		return TaskStatusTodo, nil
	case "In Progress":
		return TaskStatusInProgress, nil
	case "In Review":
		return TaskStatusInReview, nil
	case "Staging":
		return TaskStatusStaging, nil
	case "Done":
		return TaskStatusDone, nil
	}

	return TaskStatusDone, fmt.Errorf("TaskStatus=%s could not be found!", s)
}
