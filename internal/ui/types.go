package ui

import (
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/internal/ui/boards"
	"github.com/mabd-dev/gira/internal/ui/projects"
	"github.com/mabd-dev/gira/internal/ui/sprint"
)

type model struct {
	theme  theme.Theme
	width  int
	height int
	err    error

	focusStack []FocusState

	projectsModel projects.Model
	boardsModel   boards.Model
	sprintModel   sprint.Model
}
