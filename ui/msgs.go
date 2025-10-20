package ui

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/mabd-dev/gira/api"
	"github.com/mabd-dev/gira/models"
)

type fetchSprint struct {
	sprintID int
}

type fetchSprintResponse struct {
	sprint models.Sprint
	err    error
}

func (m fetchSprint) Cmd() tea.Cmd {
	return func() tea.Msg {
		// fileData, err := os.ReadFile("samples/mockApiResponses/sprintIssues.json")
		// if err != nil {
		// 	return fetchSprintResponse{err: err}
		// }

		client := api.GetClient()
		getSprintIssuesResponse, err := client.GetSprintIssues(1853)
		if err != nil {
			return fetchSprintResponse{err: err}
		}

		// var getSprintIssuesResponse api.SprintIssuesResponse
		// err = json.Unmarshal(fileData, &getSprintIssuesResponse)
		// if err != nil {
		// 	return fetchSprintResponse{err: err}
		// }

		sprint, err := models.FormatSprint(getSprintIssuesResponse)
		if err != nil {
			return fetchSprintResponse{err: err}
		}

		return fetchSprintResponse{
			sprint: sprint,
		}
	}
}
