package tasksboard

import "github.com/mabd-dev/gira/models"

type Model struct {
	tasksByStatus     map[models.TaskStatus][]models.DeveloperTask
	selectedTaskIndex int
	totalTasksCount   int
}
