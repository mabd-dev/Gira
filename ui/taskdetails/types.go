package taskdetails

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
)

type Model struct {
	viewport viewport.Model
	task     *task
	theme    theme.Theme
}

type task struct {
	DevName     string
	TaskStatus  models.TaskStatus
	Summary     string
	Description string
	StoryPoints int
}

func (m *Model) Visible() bool { return m.task != nil }
