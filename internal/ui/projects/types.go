package projects

import (
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
)

type Model struct {
	theme    theme.Theme
	projects []models.Project
	cursor   int
	offset   int
	height   int
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
