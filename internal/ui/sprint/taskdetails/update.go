package taskdetails

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit

		case "q", "esc":
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
