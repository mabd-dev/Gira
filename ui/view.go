package ui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
)

var headerBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.Border{
		Bottom:      "=",
		BottomLeft:  "=",
		BottomRight: "=",
	})

func (m model) View() string {
	header := header(m.Sprint, m.theme)
	devTabs := developersTabs(m)

	body := ""
	if m.taskDetails.Visible() {
		body = m.taskDetails.View()
	} else {
		body = m.tasksboard.View()
	}

	view := lipgloss.JoinVertical(
		lipgloss.Top,
		header,
		devTabs,
		body,
	)
	view = lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Render(view)

	return view
}

func header(sprint models.Sprint, theme theme.Theme) string {
	style := theme.Styles.Base.Foreground(theme.Colors.Foreground).Bold(true)

	boxStyle := headerBoxStyle.BorderForeground(theme.Colors.Muted)

	s := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			style.Render(sprint.Name),
			" - ",
			"6 days remaining", // TODO: calculate remaining days
		),
	)

	return boxStyle.Render(s)

}

func developersTabs(m model) string {
	devs := []string{}
	for i, developer := range m.Sprint.Developers {
		var style lipgloss.Style

		selected := m.SelectedDevIndex == i
		if selected {
			style = m.theme.Styles.Box.Foreground(m.theme.Colors.Foreground).Bold(true)
		} else {
			style = m.theme.Styles.BoxMuted.Foreground(m.theme.Colors.Muted)
		}

		devs = append(devs, style.Render(developer.Name))
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, devs...)
}
