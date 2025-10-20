// Package ui handles rendering data on terminal screen powered by bubbletea project
package ui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/ui/taskdetails"
	"github.com/mabd-dev/gira/ui/tasksboard"
)

func Render() error {
	colors, err := theme.CreateColors("catppuccin-mocha")
	if err != nil {
		return err
	}

	theme := theme.Theme{
		Colors: colors,
		Styles: theme.CreateStyles(colors),
	}

	tasksBoard := tasksboard.New(theme)
	taskDetails := taskdetails.New(theme)

	m := model{
		theme:   theme,
		loading: true,
		//Sprint:      sprint,
		tasksboard:  tasksBoard,
		taskDetails: taskDetails,
	}

	p := tea.NewProgram(m, tea.WithOutput(os.Stdout), tea.WithAltScreen())
	_, err = p.Run()
	return err

}

func (m model) Init() tea.Cmd {
	return fetchSprint{sprintID: 1122}.Cmd()
}
