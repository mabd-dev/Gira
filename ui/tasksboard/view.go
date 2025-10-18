package tasksboard

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
)

func (m *Model) View() string {
	if !hasAnyTask(m.tasksByStatus) {
		return m.theme.Styles.Base.Render("(No assigned tasks)\n")
	}

	body := ""
	taskIndex := 0

	for _, status := range models.TasksInOrder {
		taskStatusStr := m.theme.Styles.Muted.Render(string(status))
		body = lipgloss.JoinVertical(lipgloss.Left, body, taskStatusStr)

		tasks := m.tasksByStatus[status]
		if len(tasks) == 0 {
			s := m.theme.Styles.Muted.Render("  (No tasks)\n")
			body = lipgloss.JoinVertical(lipgloss.Left, body, s)
			continue
		}

		for _, task := range tasks {
			isSelected := m.selectedTaskIndex == taskIndex
			taskStr := renderTask(task, isSelected, m.theme)
			body = lipgloss.JoinVertical(lipgloss.Left, body, taskStr)

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

func renderTask(
	task models.DeveloperTask,
	isSelected bool,
	theme theme.Theme,
) string {
	var style lipgloss.Style
	if isSelected {
		style = theme.Styles.Base.Foreground(theme.Colors.Foreground).Bold(true)
	} else {
		style = theme.Styles.Base
	}

	indicator := "  "
	if isSelected {
		indicator = style.Render("> ")
	}

	storyPointsStr := style.Render("- [" + strconv.Itoa(task.StoryPoints) + "] ")

	trimmedTaskSummary := strings.TrimSpace(task.Summary)

	var taskSummaryStr string
	if isSelected {
		taskSummaryStr = style.Underline(true).Render(trimmedTaskSummary)
	} else {
		taskSummaryStr = style.Render(trimmedTaskSummary)
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		indicator,
		storyPointsStr,
		taskSummaryStr,
	)
}
