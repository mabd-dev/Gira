package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		return m, nil

	case tea.KeyMsg:

		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit

		case "tab":
			if len(m.Sprint.Developers) > 0 {
				m.SelectedDevIndex = (m.SelectedDevIndex + 1) % len(m.Sprint.Developers)
				m.tasksboard.UpdateTasks(m.Sprint.Developers[m.SelectedDevIndex].TasksByStatus)
			}
		case "shift+tab":
			devsCount := len(m.Sprint.Developers)
			if devsCount > 0 {
				m.SelectedDevIndex = (m.SelectedDevIndex - 1 + devsCount) % devsCount
				m.tasksboard.UpdateTasks(m.Sprint.Developers[m.SelectedDevIndex].TasksByStatus)
			}
		}

	}

	var cmd tea.Cmd
	m.tasksboard, cmd = m.tasksboard.Update(msg)
	return m, cmd

}
