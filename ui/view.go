package ui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/mabd-dev/gira/models"
)

var (
	box = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#89b4fa"))

	mutedBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#313244"))
)

func (m model) View() string {
	header := header(m.Sprint)
	devTabs := developersTabs(m)

	tasksBoard := m.tasksboard.View()

	return lipgloss.JoinVertical(
		lipgloss.Top,
		header,
		devTabs,
		tasksBoard,
	)
}

func header(sprint models.Sprint) string {
	style := lipgloss.NewStyle()

	return lipgloss.JoinHorizontal(
		lipgloss.Left,
		style.Render(sprint.Name),
		" - ",
		"Active (6 days remaining)",
	)
}

func developersTabs(m model) string {

	devs := []string{}
	for i, developer := range m.Sprint.Developers {
		var style lipgloss.Style
		if m.SelectedDevIndex == i {
			style = box
		} else {
			style = mutedBox
		}

		devs = append(devs, style.Render(developer.Name))
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, devs...)
}
