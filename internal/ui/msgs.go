package ui

import (
	"encoding/json"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/mabd-dev/gira/api"
	"github.com/mabd-dev/gira/models"
)

type fetchSprintResponse struct {
	sprint models.Sprint
	err    error
}

func fetchSprintCmd(sprintID int) tea.Cmd {
	return func() tea.Msg {

		// client := api.GetClient()
		// getSprintIssuesResponse, err := client.GetSprintIssues(sprintID)
		// if err != nil {
		// 	return fetchSprintResponse{err: err}
		// }

		fileData, err := os.ReadFile("samples/mockApiResponses/sprintIssues.json")
		if err != nil {
			return fetchSprintResponse{err: err}
		}

		var getSprintIssuesResponse api.SprintIssuesResponse
		err = json.Unmarshal(fileData, &getSprintIssuesResponse)
		if err != nil {
			return fetchSprintResponse{err: err}
		}

		sprint, err := models.FormatSprint(getSprintIssuesResponse)
		if err != nil {
			return fetchSprintResponse{err: err}
		}

		return fetchSprintResponse{
			sprint: sprint,
		}
	}
}
