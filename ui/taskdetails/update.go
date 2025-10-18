package taskdetails

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/internal/logger"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.task = nil
			return m, nil
		}
	case tea.WindowSizeMsg:
		// Window size is handled by parent, don't propagate to viewport
		return m, nil
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}
