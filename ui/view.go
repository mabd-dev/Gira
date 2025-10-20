package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
	"github.com/mabd-dev/gira/ui/common"
	"github.com/mabd-dev/gira/ui/taskdetails"
	"github.com/mabd-dev/gira/ui/tasksboard"
)

var headerBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.Border{
		Bottom:      "=",
		BottomLeft:  "=",
		BottomRight: "=",
	})

func (m model) View() string {
	if m.loading {
		return "Fetching sprint data ..."
	}

	header := header(m.Sprint, m.theme)
	footer := footer(m)

	devTabs := developersTabs(m)

	// Calculate heights
	headerHeight := lipgloss.Height(header)
	footerHeight := lipgloss.Height(footer)
	devTabsHeight := lipgloss.Height(devTabs)
	availableHeight := m.height - headerHeight - footerHeight - devTabsHeight

	body := ""
	if m.taskDetails.Visible() {
		m.taskDetails.UpdateSize(m.width, availableHeight)
		body = m.taskDetails.View()
	} else {
		body = m.tasksboard.View()
	}
	body = lipgloss.NewStyle().
		Height(availableHeight).
		MaxHeight(availableHeight).
		Render(body)

	view := lipgloss.JoinVertical(
		lipgloss.Top,
		header,
		devTabs,
		body,
		footer,
	)
	view = lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Render(view)

	return view
}

func header(sprint models.Sprint, theme theme.Theme) string {
	style := theme.Styles.Base.Foreground(theme.Colors.Foreground).Bold(true)

	boxStyle := headerBoxStyle.BorderForeground(theme.Colors.Muted)

	s := lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.JoinHorizontal(
			lipgloss.Left,
			style.Render(sprint.Name),
			" - ",
			"6 days remaining", // TODO: calculate remaining days
		),
	)

	return boxStyle.Render(s)

}

func developersTabs(m model) string {
	devs := []string{}
	for i, developer := range m.Sprint.Developers {
		var style lipgloss.Style

		selected := m.SelectedDevIndex == i
		if selected {
			style = m.theme.Styles.Box.Foreground(m.theme.Colors.Foreground).Bold(true)
		} else {
			style = m.theme.Styles.BoxMuted.Foreground(m.theme.Colors.Muted)
		}

		devs = append(devs, style.Render(developer.Name))
	}

	return lipgloss.JoinHorizontal(lipgloss.Left, devs...)
}

func footer(m model) string {
	if m.taskDetails.Visible() {
		return renderKeybindings(taskdetails.Keybindings, m.theme)
	} else {
		return renderKeybindings(tasksboard.Keybindings, m.theme)
	}
}

func renderKeybindings(
	keybindings []common.Keybinding,
	theme theme.Theme,
) string {
	kbStyle := theme.Styles.Base.Foreground(theme.Colors.Foreground)
	mutedStyle := theme.Styles.Muted

	var sb strings.Builder
	for i, kb := range keybindings {
		sb.WriteString(mutedStyle.Render(kb.ShortDesc))
		sb.WriteString(mutedStyle.Render(": "))
		sb.WriteString(kbStyle.Render(kb.Key))

		if i < len(keybindings)-1 {
			sb.WriteString(mutedStyle.Render(" | "))
		}
	}

	return theme.Styles.Muted.Render(sb.String())
}
