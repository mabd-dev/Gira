package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/internal/ui/boards"
	"github.com/mabd-dev/gira/internal/ui/projects"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	switch m.currentFocusState() {
	case FocusProjects:
		m.projectsModel, cmd = m.projectsModel.Update(msg)
	case FocusBoards:
		m.boardsModel, cmd = m.boardsModel.Update(msg)
	case FocusSprints:
		break
	case FocusActiveSprint:
		m.sprintModel, cmd = m.sprintModel.Update(msg)
	}

	if cmd != nil {
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		return m, cmd

	case projects.ProjectSelectedMsg:
		m.boardsModel.SetProjectID(msg.Project.ID)
		m.pushFocus(FocusBoards)
		return m, m.boardsModel.Init()

	case boards.BoardSelectedMsg:
		m.sprintModel.SetBoardID(msg.Board.ID)
		m.pushFocus(FocusActiveSprint)
		return m, m.sprintModel.Init()

	case tea.KeyMsg:

		switch msg.String() {
		case "q", "esc", "ctrl+c":
			if len(m.focusStack) == 1 {
				return m, tea.Quit
			}

			m.popFocus()
			return m, nil
		}
	}

	return m, cmd
}
