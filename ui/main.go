package ui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
	"github.com/mabd-dev/gira/ui/tasksboard"
)

func Render(sprint models.Sprint) error {
	colors, err := theme.CreateColors("catppuccin-mocha")
	if err != nil {
		return err
	}

	theme := theme.Theme{
		Colors: colors,
		Styles: theme.CreateStyles(colors),
	}

	var tasksByStatus map[models.TaskStatus][]models.DeveloperTask
	if len(sprint.Developers) > 0 {
		tasksByStatus = sprint.Developers[0].TasksByStatus
	} else {
		tasksByStatus = make(map[models.TaskStatus][]models.DeveloperTask)
	}

	tasksBoard := tasksboard.New(tasksByStatus, theme)

	m := model{
		theme:      theme,
		Sprint:     sprint,
		tasksboard: tasksBoard,
	}

	p := tea.NewProgram(m, tea.WithOutput(os.Stdout), tea.WithAltScreen())
	_, err = p.Run()
	return err

}

func (m model) Init() tea.Cmd {
	return nil
}
