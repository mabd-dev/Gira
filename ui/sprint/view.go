package sprint

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {

	if m.loading {
		loadingStyle := m.theme.Styles.Base.Foreground(m.theme.Colors.Info)
		return loadingStyle.Render("Loading projects...")
	}

	if m.err != nil {
		errorStyle := m.theme.Styles.Base.Foreground(m.theme.Colors.Error)
		return errorStyle.Render(fmt.Sprintf("Failed to load projects: %s", m.err.Error()))
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.sprint.ID,
		m.sprint.Name,
		m.sprint.Goal,
		m.sprint.StartDate,
		m.sprint.EndDate,
	)
}
