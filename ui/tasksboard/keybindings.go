package tasksboard

import "github.com/mabd-dev/gira/ui/common"

var Keybindings = []common.Keybinding{
	{
		Key:         "↑/↓ (j/k)",
		Description: "Navigate up and down (or j/k)",
		ShortDesc:   "Navigate",
	},
	{
		Key:         "<tab>",
		Description: "Next developer. (shit+tab for previous dev)",
		ShortDesc:   "Next Dev",
	},
	{
		Key:         "<enter>",
		Description: "Open task details",
		ShortDesc:   "Details",
	},
	// {
	// 	Key:         "/",
	// 	Description: "Filter by repo/branch name",
	// 	ShortDesc:   "Filter",
	// },
	{
		Key:         "q/esc",
		Description: "Quit",
		ShortDesc:   "Quit",
	},
}
