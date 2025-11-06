package taskdetails

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
)

func New(theme theme.Theme) Model {
	return Model{
		task:     nil,
		theme:    theme,
		viewport: viewport.New(100, 20),
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
	m.viewport.SetYOffset(0)
	m.viewport.SetContent(m.generateContent())
}

func (m *Model) UpdateSize(width int, height int) {
	m.viewport.Width = width
	m.viewport.Height = height
	m.viewport.SetContent(m.generateContent())
}

func (m *Model) Init() tea.Cmd { return nil }
