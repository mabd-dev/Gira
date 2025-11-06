package sprint

import (
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/internal/ui/sprint/taskdetails"
	"github.com/mabd-dev/gira/internal/ui/sprint/tasksboard"
	"github.com/mabd-dev/gira/models"
)

type Model struct {
	loading bool
	width   int
	height  int
	theme   theme.Theme
	boardID string

	sprint models.Sprint
	err    error

	SelectedDevIndex int
	tasksboardModel  tasksboard.Model
	taskDetailsModel taskdetails.Model
}

type fetchActiveSprintResponse struct {
	sprintID int
	err      error
}

type fetchActiveSprintIssuesResponse struct {
	sprint models.Sprint
	err    error
}
