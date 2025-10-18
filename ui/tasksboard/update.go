package tasksboard

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down": //move one item down
			if m.totalTasksCount > 0 {
				m.selectedTaskIndex = (m.selectedTaskIndex + 1) % m.totalTasksCount
			}
			return m, nil

		case "k", "up": //move one item up
			if m.totalTasksCount > 0 {
				m.selectedTaskIndex = (m.selectedTaskIndex - 1 + m.totalTasksCount) % m.totalTasksCount
			}
			return m, nil
		}
	}

	return m, nil
}
