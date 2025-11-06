package sprint

import (
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
)

type Model struct {
	loading bool
	theme   theme.Theme
	boardID string

	sprint models.Sprint
	err    error
}

type fetchActiveSprintResponse struct {
	sprintID int
	err      error
}

type fetchActiveSprintIssuesResponse struct {
	sprint models.Sprint
	err    error
}
