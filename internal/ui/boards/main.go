package boards

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
	m.loading = true
	m.boards = []models.Board{}
	return fetchBoardsCmd(m.projectID)
}

func (m *Model) SetProjectID(id string) {
	m.projectID = id
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
