package ui

import (
	"github.com/mabd-dev/gira/models"
	"github.com/mabd-dev/gira/ui/tasksboard"
)

type model struct {
	Sprint           models.Sprint
	SelectedDevIndex int
	width            int
	height           int

	tasksboard tasksboard.Model
}
