// Package tassboard is reponsible for displaying tasks per developer
package tasksboard

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/models"
)

func New(tasksByStatus map[models.TaskStatus][]models.DeveloperTask) Model {
	totalTasksCount := 0
	for _, tasks := range tasksByStatus {
		totalTasksCount += len(tasks)
	}

	return Model{
		tasksByStatus:     tasksByStatus,
		selectedTaskIndex: 0,
		totalTasksCount:   totalTasksCount,
	}
}

func (m *Model) Init() tea.Cmd { return nil }

func (m *Model) UpdateTasks(
	tasksByStatus map[models.TaskStatus][]models.DeveloperTask,
) {
	m.tasksByStatus = tasksByStatus

	if m.totalTasksCount > 0 {
		m.selectedTaskIndex = 0
	} else {
		m.selectedTaskIndex = -1
	}
}
