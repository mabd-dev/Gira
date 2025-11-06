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

	for i, project := range m.projects {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		line := fmt.Sprintf("%s %s (%s)\n", cursor, project.Name, project.ProjectTypeKey)
		b.WriteString(line)
	}

	return b.String()
}
