package boards

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	if m.loading {
		return "Loading..."
	}

	if m.err != nil {
		return fmt.Sprintf("Failed to load boards: %s", m.err.Error())
	}

	if len(m.boards) == 0 {
		return "No boards found"
	}

	var b strings.Builder
	b.WriteString("Boards:\n\n")

	// Calculate visible range
	start := m.offset
	end := m.offset + m.height
	if end > len(m.boards) {
		end = len(m.boards)
	}

	// Show indicator if there are items above
	if m.offset > 0 {
		b.WriteString(fmt.Sprintf("  ... (%d more above)\n", m.offset))
	}

	// Show visible items
	for i := start; i < end; i++ {
		board := m.boards[i]
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		privacy := "public"
		if board.IsPrivate {
			privacy = "private"
		}

		line := fmt.Sprintf("%s %s (%s)\n", cursor, board.Name, privacy)
		b.WriteString(line)
	}

	// Show indicator if there are items below
	if end < len(m.boards) {
		b.WriteString(fmt.Sprintf("  ... (%d more below)\n", len(m.boards)-end))
	}

	return b.String()
}
