package tasksboard

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/internal/logger"
	"github.com/mabd-dev/gira/models"
)

type NoMsg struct{}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			status, taskIndex, err := findSelectedTask(m.tasksByStatus, m.selectedTaskIndex)
			if err != nil {
				logger.Debug(err.Error())
				return m, nil
			}
			return m, func() tea.Msg {
				return TaskSelectedMsg{
					Status:    status,
					TaskIndex: taskIndex,
				}
			}
		case "j", "down": //move one item down
			if m.totalTasksCount > 0 {
				m.selectedTaskIndex = (m.selectedTaskIndex + 1) % m.totalTasksCount
			}
			return m, nil

		case "k", "up": //move one item up
			if m.totalTasksCount > 0 {
				m.selectedTaskIndex = (m.selectedTaskIndex - 1 + m.totalTasksCount) % m.totalTasksCount
			}
			return m, nil
		}
	}

	return m, nil
}

func findSelectedTask(
	tasksByStatus map[models.TaskStatus][]models.DeveloperTask,
	selectedTaskIndex int,
) (models.TaskStatus, int, error) {
	currTaskIndex := 0

	for _, status := range models.TaskStatusInOrder {
		tasks := tasksByStatus[status]

		for index := range tasks {
			selected := currTaskIndex == selectedTaskIndex
			if selected {
				return status, index, nil
			}

			currTaskIndex++
		}
	}

	return models.TaskStatusTodo, 0, fmt.Errorf("not able to find selected task!")
}
