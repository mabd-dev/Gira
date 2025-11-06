package projects

import (
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/api"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
)

func New(t theme.Theme) Model {
	return Model{
		theme:   t,
		loading: true,
		height:  20, // Default height, will be updated with WindowSizeMsg
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

		// Sort projects by ID
		sort.Slice(projects, func(i, j int) bool {
			return projects[i].ID < projects[j].ID
		})

		return fetchProjectsResponse{
			projects: projects,
		}
	}
}
