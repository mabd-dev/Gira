package boards

import (
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
)

type Model struct {
	theme     theme.Theme
	projectID string
	boards    []models.Board
	cursor    int
	offset    int
	height    int
	loading   bool
	err       error
}

type BoardSelectedMsg struct {
	Board models.Board
}

type fetchBoardsResponse struct {
	boards []models.Board
	err    error
}
