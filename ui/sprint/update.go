package sprint

import tea "github.com/charmbracelet/bubbletea"

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case fetchActiveSprintResponse:
		if msg.err != nil {
			m.loading = false
			m.err = msg.err
			return m, nil
		}
		return m, fetchActiveSprintIssues(msg.sprintID)
	case fetchActiveSprintIssuesResponse:
		if msg.err != nil {
			m.loading = false
			m.err = msg.err
			return m, nil
		}
		m.sprint = msg.sprint
		m.loading = false
		return m, nil
	}
	return m, nil
}
