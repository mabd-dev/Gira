package taskdetails

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
)

func New(theme theme.Theme) Model {
	return Model{
		task:  nil,
		theme: theme,
	}
}

func (m *Model) Show(
	devName string,
	taskStatus models.TaskStatus,
	summary string,
	description string,
	storyPoints int,
) {
	task := task{
		DevName:     devName,
		TaskStatus:  taskStatus,
		Summary:     summary,
		Description: description,
		StoryPoints: storyPoints,
	}
	m.task = &task
}

func (m *Model) Init() tea.Cmd { return nil }
