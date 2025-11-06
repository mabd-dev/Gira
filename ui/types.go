package ui

import (
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
	"github.com/mabd-dev/gira/ui/boards"
	"github.com/mabd-dev/gira/ui/projects"
	"github.com/mabd-dev/gira/ui/sprint"
	"github.com/mabd-dev/gira/ui/taskdetails"
	"github.com/mabd-dev/gira/ui/tasksboard"
)

type model struct {
	theme  theme.Theme
	width  int
	height int
	err    error

	Sprint           models.Sprint
	SelectedDevIndex int

	focusStack []FocusState

	projectsModel    projects.Model
	boardsModel      boards.Model
	sprintModel      sprint.Model
	tasksboardModel  tasksboard.Model
	taskDetailsModel taskdetails.Model
}
