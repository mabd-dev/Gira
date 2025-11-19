package sprint

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/internal/ui/common"
	"github.com/mabd-dev/gira/internal/ui/sprint/tasksboard"
	"github.com/mabd-dev/gira/models"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {

	var cmd tea.Cmd

	if m.taskDetailsModel.Visible() {
		m.taskDetailsModel, cmd = m.taskDetailsModel.Update(msg)
		if cmd == nil {
			cmd = common.ExitCmd
		}
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		m.taskDetailsModel, cmd = m.taskDetailsModel.Update(msg)
		return m, cmd

	case fetchActiveSprintResponse:
		if msg.err != nil {
			m.loading = false
			m.err = msg.err
			m.activeSprintID = -1
			return m, nil
		}
		m.activeSprintID = msg.sprintID
		return m, fetchActiveSprintIssues(m.activeSprintID)
	case fetchActiveSprintIssuesResponse:
		if msg.err != nil {
			m.loading = false
			m.err = msg.err
			return m, nil
		}
		m.sprint = msg.sprint
		m.loading = false

		var tasksByStatus map[models.TaskStatus][]models.DeveloperTask
		if len(msg.sprint.Developers) > 0 {
			index := 0
			if m.SelectedDevIndex <= len(msg.sprint.Developers)-1 {
				index = m.SelectedDevIndex
			}
			tasksByStatus = msg.sprint.Developers[index].TasksByStatus
		} else {
			tasksByStatus = make(map[models.TaskStatus][]models.DeveloperTask)
		}
		m.tasksboardModel.UpdateTasks(tasksByStatus)
		return m, nil

	case tasksboard.TaskSelectedMsg:
		dev := m.sprint.Developers[m.SelectedDevIndex]
		task := m.sprint.Developers[m.SelectedDevIndex].TasksByStatus[msg.Status][msg.TaskIndex]

		m.taskDetailsModel.Show(
			dev.Name,
			msg.Status,
			task.Summary,
			task.Description,
			task.StoryPoints,
		)
		return m, common.ExitCmd

	case tea.KeyMsg:
		switch msg.String() {
		case "r":
			if m.activeSprintID > 0 {
				m.loading = true
				return m, fetchActiveSprintIssues(m.activeSprintID)
			}
			return m, nil
		case "tab":
			if len(m.sprint.Developers) > 0 {
				m.SelectedDevIndex = (m.SelectedDevIndex + 1) % len(m.sprint.Developers)
				m.tasksboardModel.UpdateTasks(m.sprint.Developers[m.SelectedDevIndex].TasksByStatus)
			}

		case "shift+tab":
			devsCount := len(m.sprint.Developers)
			if devsCount > 0 {
				m.SelectedDevIndex = (m.SelectedDevIndex - 1 + devsCount) % devsCount
				m.tasksboardModel.UpdateTasks(m.sprint.Developers[m.SelectedDevIndex].TasksByStatus)
			}
		}
	}

	m.tasksboardModel, cmd = m.tasksboardModel.Update(msg)
	return m, cmd
}
