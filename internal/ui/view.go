package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/internal/ui/common"
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
	footer := renderKeybindings(ErrorFetchingSprintKeybindings, m.theme)

	headerHeight := lipgloss.Height(header)
	footerHeight := lipgloss.Height(footer)

	availableHeight := m.height - headerHeight - footerHeight

	body := lipgloss.NewStyle().
		Height(availableHeight).
		Render("")

	return lipgloss.JoinVertical(
		lipgloss.Top,
		header,
		body,
		footer,
	)
}

func renderKeybindings(
	keybindings []common.Keybinding,
	theme theme.Theme,
) string {
	kbStyle := theme.Styles.Base.Foreground(theme.Colors.Foreground)
	mutedStyle := theme.Styles.Muted

	var sb strings.Builder
	for i, kb := range keybindings {
		sb.WriteString(mutedStyle.Render(kb.ShortDesc))
		sb.WriteString(mutedStyle.Render(": "))
		sb.WriteString(kbStyle.Render(kb.Key))

		if i < len(keybindings)-1 {
			sb.WriteString(mutedStyle.Render(" | "))
		}
	}

	return theme.Styles.Muted.Render(sb.String())
}
