package ui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {

	switch m.currentFocusState() {
	case FocusProjects:
		return m.projectsModel.View()
	case FocusBoards:
		return m.boardsModel.View()
	case FocusSprints:
		break
	case FocusActiveSprint:
		return m.sprintModel.View()
	}

	if m.err != nil {
		return renderErrorFetching(m)
	}

	return ""
}

func renderErrorFetching(m model) string {
	header := lipgloss.JoinVertical(
		lipgloss.Top,
		m.theme.Styles.Base.Foreground(m.theme.Colors.Error).Bold(true).Render("Error"),
		m.theme.Styles.Base.Foreground(m.theme.Colors.Foreground).Render(m.err.Error()),
	)

	headerHeight := lipgloss.Height(header)

	availableHeight := m.height - headerHeight

	body := lipgloss.NewStyle().
		Height(availableHeight).
		Render("")

	return lipgloss.JoinVertical(
		lipgloss.Top,
		header,
		body,
	)
}
