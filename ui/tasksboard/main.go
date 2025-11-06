// Package tasksboard is reponsible for displaying tasks per developer
package tasksboard

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
)

func New(
	theme theme.Theme,
) Model {
	return Model{
		tasksByStatus:     map[models.TaskStatus][]models.DeveloperTask{},
		theme:             theme,
		selectedTaskIndex: 0,
		totalTasksCount:   0,
	}
}

func (m *Model) Init() tea.Cmd { return nil }

func (m *Model) UpdateTasks(
	tasksByStatus map[models.TaskStatus][]models.DeveloperTask,
) {
	m.tasksByStatus = tasksByStatus
	m.totalTasksCount = totalTasksCount(tasksByStatus)

	if m.totalTasksCount > 0 {
		m.selectedTaskIndex = 0
	} else {
		m.selectedTaskIndex = -1
	}
}

func totalTasksCount(
	tasksByStatus map[models.TaskStatus][]models.DeveloperTask,
) int {
	totalTasksCount := 0
	for _, tasks := range tasksByStatus {
		totalTasksCount += len(tasks)
	}
	return totalTasksCount
}
