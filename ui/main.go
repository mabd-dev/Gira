package ui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/models"
	"github.com/mabd-dev/gira/ui/tasksboard"
)

func Render(sprint models.Sprint) error {
	var tasksBoard tasksboard.Model
	if len(sprint.Developers) > 0 {
		tasksBoard = tasksboard.New(sprint.Developers[0].TasksByStatus)
	} else {
		tasksBoard = tasksboard.New(make(map[models.TaskStatus][]models.DeveloperTask))
	}

	m := model{
		Sprint:     sprint,
		tasksboard: tasksBoard,
	}

	p := tea.NewProgram(m, tea.WithOutput(os.Stdout), tea.WithAltScreen())
	_, err := p.Run()
	return err

}

func (m model) Init() tea.Cmd {
	return nil
}
