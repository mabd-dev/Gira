package projects

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case fetchProjectsResponse:
		if msg.err != nil {
			m.err = msg.err
			m.loading = false
			return m, nil
		}
		m.projects = msg.projects
		m.loading = false
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.projects)-1 {
				m.cursor++
			}
		case "enter":
			if len(m.projects) > 0 && m.cursor < len(m.projects) {
				return m, func() tea.Msg {
					return ProjectSelectedMsg{
						Project: m.projects[m.cursor],
					}
				}
			}
		}
	}

	return m, nil
}
