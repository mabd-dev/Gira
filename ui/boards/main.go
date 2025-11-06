package boards

import (
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/api"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
)

func New(theme theme.Theme, projectID string) Model {
	return Model{
		projectID: projectID,
		loading:   true,
		height:    20, // Default height, will be updated with WindowSizeMsg
	}
}

func (m *Model) Init() tea.Cmd {
	return fetchBoardsCmd(m.projectID)
}

func fetchBoardsCmd(projectID string) tea.Cmd {
	return func() tea.Msg {
		client := api.GetClient()
		boardsResponse, err := client.GetBoards(projectID)
		if err != nil {
			return fetchBoardsResponse{err: err}
		}

		boards, err := models.FormatBoards(boardsResponse)
		if err != nil {
			return fetchBoardsResponse{err: err}
		}

		// Sort boards by ID
		sort.Slice(boards, func(i, j int) bool {
			return boards[i].ID < boards[j].ID
		})

		return fetchBoardsResponse{
			boards: boards,
		}
	}
}
