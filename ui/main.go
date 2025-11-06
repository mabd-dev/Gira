// Package ui handles rendering data on terminal screen powered by bubbletea project
package ui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/ui/boards"
	"github.com/mabd-dev/gira/ui/projects"
	"github.com/mabd-dev/gira/ui/sprint"
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

	m := model{
		theme:            theme,
		focusStack:       []FocusState{FocusProjects},
		projectsModel:    projects.New(theme),
		sprintModel:      sprint.New(theme),
		boardsModel:      boards.New(theme),
		tasksboardModel:  tasksboard.New(theme),
		taskDetailsModel: taskdetails.New(theme),
	}

	p := tea.NewProgram(m, tea.WithOutput(os.Stdout), tea.WithAltScreen())
	_, err = p.Run()
	return err

}

func (m model) Init() tea.Cmd {
	return m.projectsModel.Init()
}
