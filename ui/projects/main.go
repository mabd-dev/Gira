package projects

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/api"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
)

func New(theme theme.Theme) Model {
	return Model{
		loading: true,
	}
}

func (m *Model) Init() tea.Cmd {
	return fetchProjectsCmd()
}

func fetchProjectsCmd() tea.Cmd {
	return func() tea.Msg {
		client := api.GetClient()
		projectsResponse, err := client.GetProjects()
		if err != nil {
			return fetchProjectsResponse{err: err}
		}

		projects, err := models.FormatProject(projectsResponse)
		if err != nil {
			return fetchProjectsResponse{err: err}
		}

		return fetchProjectsResponse{
			projects: projects,
		}
	}
}
