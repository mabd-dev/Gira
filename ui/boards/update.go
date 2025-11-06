package boards

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case fetchBoardsResponse:
		if msg.err != nil {
			m.err = msg.err
			m.loading = false
			return m, nil
		}
		m.boards = msg.boards
		m.loading = false
		return m, nil

	case tea.WindowSizeMsg:
		// Reserve space for header and footer (adjust as needed)
		m.height = msg.Height - 4
		if m.height < 5 {
			m.height = 5
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
				// Scroll up if cursor moves above visible area
				if m.cursor < m.offset {
					m.offset = m.cursor
				}
			}
		case "down", "j":
			if m.cursor < len(m.boards)-1 {
				m.cursor++
				// Scroll down if cursor moves below visible area
				if m.cursor >= m.offset+m.height {
					m.offset = m.cursor - m.height + 1
				}
			}
		case "enter":
			if len(m.boards) > 0 && m.cursor < len(m.boards) {
				return m, func() tea.Msg {
					return BoardSelectedMsg{
						Board: m.boards[m.cursor],
					}
				}
			}
		}
	}

	return m, nil
}
