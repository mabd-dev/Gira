package projects

import (
	"github.com/mabd-dev/gira/models"
)

type Model struct {
	projects []models.Project
	cursor   int
	loading  bool
	err      error
}

type ProjectSelectedMsg struct {
	Project models.Project
}

type fetchProjectsResponse struct {
	projects []models.Project
	err      error
}
