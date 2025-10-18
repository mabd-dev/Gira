package taskdetails

import (
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	sectionStyle := lipgloss.NewStyle().MarginBottom(1)
	header := sectionStyle.Render(renderHeader(m))
	taskSummary := sectionStyle.Render(renderTaskSummary(m))

	taskDescription := sectionStyle.Render(renderTaskDescription(m))

	return lipgloss.JoinVertical(
		lipgloss.Top,
		"<- (back)\n",
		header,
		taskSummary,
		taskDescription,
	)
}

func renderHeader(m Model) string {
	fgStyle := m.theme.Styles.Base.Foreground(m.theme.Colors.Foreground)

	return lipgloss.JoinHorizontal(
		lipgloss.Center,

		m.theme.Styles.Muted.Render("status:"),
		fgStyle.Render(string(m.task.TaskStatus)),

		m.theme.Styles.Muted.Render(" | "),
		m.theme.Styles.Muted.Render("sp:"),
		fgStyle.Render(strconv.Itoa(m.task.StoryPoints)),
	)
}

func renderTaskSummary(m Model) string {
	fgStyle := m.theme.Styles.Base.Foreground(m.theme.Colors.Foreground)

	return lipgloss.JoinVertical(
		lipgloss.Left,

		m.theme.Styles.Muted.Render("task:"),
		fgStyle.Render(m.task.Summary),
	)
}

func renderTaskDescription(m Model) string {
	if len(m.task.Description) == 0 {
		return ""
	}

	fgStyle := m.theme.Styles.Base.Foreground(m.theme.Colors.Foreground)

	return lipgloss.JoinVertical(
		lipgloss.Left,

		m.theme.Styles.Muted.Render("description:"),
		fgStyle.Render(m.task.Description),
	)
}
