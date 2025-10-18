package tasksboard

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mabd-dev/gira/models"
)

func (m *Model) View() string {
	if !hasAnyTask(m.tasksByStatus) {
		return "(No assigned tasks)\n"
	}

	body := ""
	var sb strings.Builder
	taskIndex := 0

	for _, status := range models.TasksInOrder {
		body = lipgloss.JoinVertical(lipgloss.Left, body, string(status))

		tasks := m.tasksByStatus[status]
		if len(tasks) == 0 {
			body = lipgloss.JoinVertical(lipgloss.Left, body, "  (No tasks)\n")
			continue
		}

		for _, task := range tasks {
			indicator := "  "
			if m.selectedTaskIndex == taskIndex {
				indicator = "> "
			}

			sb.WriteString(indicator)
			sb.WriteString("- [")
			sb.WriteString(strconv.Itoa(task.StoryPoints))
			sb.WriteString("] ")

			sb.WriteString(strings.TrimSpace(task.Summary))

			body = lipgloss.JoinVertical(lipgloss.Left, body, sb.String())

			sb.Reset()

			taskIndex++
		}

		body = lipgloss.JoinVertical(lipgloss.Left, body, "")
	}

	return body
}

func hasAnyTask(tasksByStatus map[models.TaskStatus][]models.DeveloperTask) bool {
	for _, status := range models.TasksInOrder {
		tasks := tasksByStatus[status]
		if len(tasks) > 0 {
			return true
		}
	}
	return false
}
