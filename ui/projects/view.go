package projects

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	if m.loading {
		return "Loading..."
	}

	if m.err != nil {
		return fmt.Sprintf("Failed to load projects: %s", m.err.Error())
	}

	if len(m.projects) == 0 {
		return "No projects found"
	}

	var b strings.Builder
	b.WriteString("Projects:\n\n")

	// Calculate visible range
	start := m.offset
	end := m.offset + m.height
	if end > len(m.projects) {
		end = len(m.projects)
	}

	// Show indicator if there are items above
	if m.offset > 0 {
		b.WriteString(fmt.Sprintf("  ... (%d more above)\n", m.offset))
	}

	// Show visible items
	for i := start; i < end; i++ {
		project := m.projects[i]
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		line := fmt.Sprintf("%s %s (%s)\n", cursor, project.Name, project.ProjectTypeKey)
		b.WriteString(line)
	}

	// Show indicator if there are items below
	if end < len(m.projects) {
		b.WriteString(fmt.Sprintf("  ... (%d more below)\n", len(m.projects)-end))
	}

	return b.String()
}
