package tasksboard

import (
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
)

type Model struct {
	theme             theme.Theme
	tasksByStatus     map[models.TaskStatus][]models.DeveloperTask
	selectedTaskIndex int
	totalTasksCount   int
}
