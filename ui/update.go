package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/internal/logger"
	"github.com/mabd-dev/gira/models"
	"github.com/mabd-dev/gira/ui/projects"
	"github.com/mabd-dev/gira/ui/tasksboard"
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
		break
	}

	if cmd != nil {
		return m, cmd
	}

	if m.taskDetailsModel.Visible() {
		m.taskDetailsModel, cmd = m.taskDetailsModel.Update(msg)
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.taskDetailsModel, cmd = m.taskDetailsModel.Update(msg)
		return m, cmd

	case projects.ProjectSelectedMsg:
		m.boardsModel.SetProjectID(msg.Project.ID)
		m.pushFocus(FocusBoards)
		return m, m.boardsModel.Init()

	case fetchSprintResponse:
		m.err = msg.err

		if msg.err != nil {
			logger.Error("error fetching sprint data", logger.StringAttr("error", msg.err.Error()))
			// m.loading = false
			return m, nil
		}
		m.Sprint = msg.sprint

		var tasksByStatus map[models.TaskStatus][]models.DeveloperTask
		if len(msg.sprint.Developers) > 0 {
			index := 0
			if m.SelectedDevIndex < len(msg.sprint.Developers)-1 {
				index = m.SelectedDevIndex
			}
			tasksByStatus = msg.sprint.Developers[index].TasksByStatus
		} else {
			tasksByStatus = make(map[models.TaskStatus][]models.DeveloperTask)
		}
		m.tasksboardModel.UpdateTasks(tasksByStatus)
		// m.loading = false
		return m, nil

	case tasksboard.TaskSelectedMsg:
		dev := m.Sprint.Developers[m.SelectedDevIndex]
		task := m.Sprint.Developers[m.SelectedDevIndex].TasksByStatus[msg.Status][msg.TaskIndex]

		m.taskDetailsModel.Show(
			dev.Name,
			msg.Status,
			task.Summary,
			task.Description,
			task.StoryPoints,
		)

	case tea.KeyMsg:

		switch msg.String() {
		case "q", "esc", "ctrl+c":
			if len(m.focusStack) == 1 {
				return m, tea.Quit
			}

			m.popFocus()
			return m, nil

		// case "r":
		// 	return m, fetchSprintCmd(1)

		case "tab":
			if len(m.Sprint.Developers) > 0 {
				m.SelectedDevIndex = (m.SelectedDevIndex + 1) % len(m.Sprint.Developers)
				m.tasksboardModel.UpdateTasks(m.Sprint.Developers[m.SelectedDevIndex].TasksByStatus)
			}
		case "shift+tab":
			devsCount := len(m.Sprint.Developers)
			if devsCount > 0 {
				m.SelectedDevIndex = (m.SelectedDevIndex - 1 + devsCount) % devsCount
				m.tasksboardModel.UpdateTasks(m.Sprint.Developers[m.SelectedDevIndex].TasksByStatus)
			}
		}

	}

	m.tasksboardModel, cmd = m.tasksboardModel.Update(msg)
	return m, cmd
}
