package projects

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mabd-dev/gira/internal/theme"
)

func New(theme theme.Theme) Model {
	return Model{}
}

func (m *Model) Init() tea.Cmd { return nil }
