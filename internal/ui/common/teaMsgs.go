package common

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Exit struct{}

var ExitCmd tea.Cmd = func() tea.Msg {
	return Exit{}
}
