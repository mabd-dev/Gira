package sprint

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/api"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
	"github.com/mabd-dev/gira/ui/taskdetails"
	"github.com/mabd-dev/gira/ui/tasksboard"
)

func New(t theme.Theme) Model {
	return Model{
		theme:   t,
		loading: true,

		tasksboardModel:  tasksboard.New(t),
		taskDetailsModel: taskdetails.New(t),
	}
}

func (m *Model) Init() tea.Cmd {
	m.loading = true
	m.sprint = models.Sprint{}
	m.tasksboardModel.UpdateTasks(map[models.TaskStatus][]models.DeveloperTask{})
	return fetchActiveSprint(m.boardID)
}

func (m *Model) SetBoardID(id int) {
	m.boardID = strconv.Itoa(id)
}

func fetchActiveSprint(boardID string) tea.Cmd {
	return func() tea.Msg {
		client := api.GetClient()
		activeSprintResponse, err := client.GetActiveSprint(boardID)
		if err != nil {
			return fetchActiveSprintResponse{err: err}
		}

		return fetchActiveSprintResponse{sprintID: activeSprintResponse.ID}
	}
}

func fetchActiveSprintIssues(sprintID int) tea.Cmd {
	return func() tea.Msg {
		client := api.GetClient()
		activeSprintIssuesResponse, err := client.GetSprintIssues(sprintID)
		if err != nil {
			return fetchActiveSprintIssuesResponse{err: err}
		}

		sprint, err := models.FormatSprint(activeSprintIssuesResponse)

		return fetchActiveSprintIssuesResponse{sprint: sprint, err: err}
	}
}
