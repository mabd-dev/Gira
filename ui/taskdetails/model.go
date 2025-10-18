package taskdetails

import (
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
)

type Model struct {
	task  *task
	theme theme.Theme
}

type task struct {
	DevName     string
	TaskStatus  models.TaskStatus
	Summary     string
	Description string
	StoryPoints int
}

func (m *Model) Visible() bool { return m.task != nil }
