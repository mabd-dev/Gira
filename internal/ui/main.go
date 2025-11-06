// Package ui handles rendering data on terminal screen powered by bubbletea project
package ui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/internal/ui/boards"
	"github.com/mabd-dev/gira/internal/ui/projects"
	"github.com/mabd-dev/gira/internal/ui/sprint"
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
		theme:         theme,
		focusStack:    []FocusState{FocusProjects},
		projectsModel: projects.New(theme),
		sprintModel:   sprint.New(theme),
		boardsModel:   boards.New(theme),
	}

	p := tea.NewProgram(m, tea.WithOutput(os.Stdout), tea.WithAltScreen())
	_, err = p.Run()
	return err

}

func (m model) Init() tea.Cmd {
	return m.projectsModel.Init()
}
