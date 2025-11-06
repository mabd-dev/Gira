package boards

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.loading {
		loadingStyle := m.theme.Styles.Base.Foreground(m.theme.Colors.Info)
		return loadingStyle.Render("Loading boards...")
	}

	if m.err != nil {
		errorStyle := m.theme.Styles.Base.Foreground(m.theme.Colors.Error)
		return errorStyle.Render(fmt.Sprintf("Failed to load boards: %s", m.err.Error()))
	}

	if len(m.boards) == 0 {
		emptyStyle := m.theme.Styles.Muted
		return emptyStyle.Render("No boards found")
	}

	var b strings.Builder

	// Header
	headerStyle := m.theme.Styles.TableHeader
	header := headerStyle.Render("🎯 Boards")
	b.WriteString(header + "\n\n")

	// Calculate visible range
	start := m.offset
	end := m.offset + m.height
	if end > len(m.boards) {
		end = len(m.boards)
	}

	// Show indicator if there are items above
	if m.offset > 0 {
		scrollStyle := m.theme.Styles.Muted
		indicator := scrollStyle.Render(fmt.Sprintf("  ▲ %d more above", m.offset))
		b.WriteString(indicator + "\n")
	}

	// Show visible items
	for i := start; i < end; i++ {
		board := m.boards[i]
		isSelected := m.cursor == i

		var line string
		if isSelected {
			selectedStyle := m.theme.Styles.Base.Foreground(m.theme.Colors.Accent).Bold(true)

			cursor := selectedStyle.Render("▶ ")
			name := selectedStyle.Underline(true).Render(board.Name)

			// Privacy indicator with color
			var privacyIcon, privacyText string
			if board.IsPrivate {
				privacyIcon = "🔒"
				privacyText = "private"
			} else {
				privacyIcon = "🌐"
				privacyText = "public"
			}
			privacyStyle := m.theme.Styles.Base.Foreground(m.theme.Colors.Info)
			privacy := privacyStyle.Render(fmt.Sprintf(" %s %s", privacyIcon, privacyText))

			line = lipgloss.JoinHorizontal(lipgloss.Left, cursor, name, privacy)
		} else {
			normalStyle := m.theme.Styles.Base
			metaStyle := m.theme.Styles.Muted

			cursor := "  "
			name := normalStyle.Render(board.Name)

			// Privacy indicator
			var privacyIcon, privacyText string
			if board.IsPrivate {
				privacyIcon = "🔒"
				privacyText = "private"
			} else {
				privacyIcon = "🌐"
				privacyText = "public"
			}
			privacy := metaStyle.Render(fmt.Sprintf(" %s %s", privacyIcon, privacyText))

			line = lipgloss.JoinHorizontal(lipgloss.Left, cursor, name, privacy)
		}

		b.WriteString(line + "\n")
	}

	// Show indicator if there are items below
	if end < len(m.boards) {
		scrollStyle := m.theme.Styles.Muted
		indicator := scrollStyle.Render(fmt.Sprintf("  ▼ %d more below", len(m.boards)-end))
		b.WriteString(indicator)
	}

	return b.String()
}
