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

	for _, status := range models.TaskStatusInOrder {
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
	for _, status := range models.TaskStatusInOrder {
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

	metadata := renderTaskMetaData(task, theme)

	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			indicator,
			storyPointsStr,
			taskSummaryStr,
		),
		metadata,
	)
}

func renderTaskMetaData(
	task models.DeveloperTask,
	theme theme.Theme,
) string {
	if len(task.Components) == 0 || len(task.FixVersions) == 0 {
		return ""
	}

	style := theme.Styles.Muted

	components := ""
	versions := ""

	if len(task.Components) > 0 {
		components = "    components: [" + strings.Join(task.Components, ",") + "]"
	}

	if len(task.FixVersions) > 0 {
		versions = "V: [" + strings.Join(task.FixVersions, ",") + "]"
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		style.Render(components),
		style.Render(" | "),
		style.Render(versions),
	) + "\n"
}
