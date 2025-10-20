package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/models"
	"github.com/mabd-dev/gira/ui/tasksboard"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	var cmd tea.Cmd

	if m.taskDetails.Visible() {
		m.taskDetails, cmd = m.taskDetails.Update(msg)
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.taskDetails, cmd = m.taskDetails.Update(msg)
		return m, cmd

	case fetchSprintResponse:
		m.Sprint = msg.sprint

		var tasksByStatus map[models.TaskStatus][]models.DeveloperTask
		if len(msg.sprint.Developers) > 0 {
			tasksByStatus = msg.sprint.Developers[0].TasksByStatus
		} else {
			tasksByStatus = make(map[models.TaskStatus][]models.DeveloperTask)
		}
		m.tasksboard.UpdateTasks(tasksByStatus)
		m.loading = false
		return m, nil

	case tasksboard.TaskSelected:
		dev := m.Sprint.Developers[m.SelectedDevIndex]
		task := m.Sprint.Developers[m.SelectedDevIndex].TasksByStatus[msg.Status][msg.TaskIndex]

		m.taskDetails.Show(
			dev.Name,
			msg.Status,
			task.Summary,
			task.Description,
			task.StoryPoints,
		)

	case tea.KeyMsg:

		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit

		case "r":
			m.loading = true
			return m, fetchSprint{sprintID: 1}.Cmd()

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

	m.tasksboard, cmd = m.tasksboard.Update(msg)
	return m, cmd
}
