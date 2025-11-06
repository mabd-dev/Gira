package ui

import "github.com/mabd-dev/gira/internal/ui/common"

var ErrorFetchingSprintKeybindings = []common.Keybinding{
	{
		Key:         "r",
		Description: "Refetch sprint tasks",
		ShortDesc:   "Refresh",
	},
	{
		Key:         "q/esc",
		Description: "Quit",
		ShortDesc:   "Quit",
	},
}
