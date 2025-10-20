package ui

import (
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
	"github.com/mabd-dev/gira/ui/taskdetails"
	"github.com/mabd-dev/gira/ui/tasksboard"
)

type model struct {
	theme   theme.Theme
	width   int
	height  int
	loading bool
	err     error

	Sprint           models.Sprint
	SelectedDevIndex int

	tasksboard  tasksboard.Model
	taskDetails taskdetails.Model
}
