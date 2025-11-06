package projects

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.loading {
		loadingStyle := m.theme.Styles.Base.Foreground(m.theme.Colors.Info)
		return loadingStyle.Render("Loading projects...")
	}

	if m.err != nil {
		errorStyle := m.theme.Styles.Base.Foreground(m.theme.Colors.Error)
		return errorStyle.Render(fmt.Sprintf("Failed to load projects: %s", m.err.Error()))
	}

	if len(m.projects) == 0 {
		emptyStyle := m.theme.Styles.Muted
		return emptyStyle.Render("No projects found")
	}

	var b strings.Builder

	// Header
	headerStyle := m.theme.Styles.TableHeader
	header := headerStyle.Render("📋 Projects")
	b.WriteString(header + "\n\n")

	// Calculate visible range
	start := m.offset
	end := m.offset + m.height
	if end > len(m.projects) {
		end = len(m.projects)
	}

	// Show indicator if there are items above
	if m.offset > 0 {
		scrollStyle := m.theme.Styles.Muted
		indicator := scrollStyle.Render(fmt.Sprintf("  ▲ %d more above", m.offset))
		b.WriteString(indicator + "\n")
	}

	// Show visible items
	for i := start; i < end; i++ {
		project := m.projects[i]
		isSelected := m.cursor == i

		var line string
		if isSelected {
			selectedStyle := m.theme.Styles.Base.Foreground(m.theme.Colors.Accent).Bold(true)
			metaStyle := m.theme.Styles.Base.Foreground(m.theme.Colors.Info)

			cursor := selectedStyle.Render("▶ ")
			name := selectedStyle.Underline(true).Render(project.Name)
			meta := metaStyle.Render(fmt.Sprintf(" (%s)", project.ProjectTypeKey))

			line = lipgloss.JoinHorizontal(lipgloss.Left, cursor, name, meta)
		} else {
			normalStyle := m.theme.Styles.Base
			metaStyle := m.theme.Styles.Muted

			cursor := "  "
			name := normalStyle.Render(project.Name)
			meta := metaStyle.Render(fmt.Sprintf(" (%s)", project.ProjectTypeKey))

			line = lipgloss.JoinHorizontal(lipgloss.Left, cursor, name, meta)
		}

		b.WriteString(line + "\n")
	}

	// Show indicator if there are items below
	if end < len(m.projects) {
		scrollStyle := m.theme.Styles.Muted
		indicator := scrollStyle.Render(fmt.Sprintf("  ▼ %d more below", len(m.projects)-end))
		b.WriteString(indicator)
	}

	return b.String()
}
