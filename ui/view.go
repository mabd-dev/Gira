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

	switch m.currentFocusState() {
	case FocusProjects:
		return m.projectsModel.View()
	case FocusBoards:
		return m.boardsModel.View()
	case FocusSprints:
		break
	case FocusActiveSprint:
		return m.sprintModel.View()
	}

	if m.err != nil {
		return renderErrorFetching(m)
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
	if m.taskDetailsModel.Visible() {
		m.taskDetailsModel.UpdateSize(m.width, availableHeight)
		body = m.taskDetailsModel.View()
	} else {
		body = m.tasksboardModel.View()
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

func renderErrorFetching(m model) string {
	header := lipgloss.JoinVertical(
		lipgloss.Top,
		m.theme.Styles.Base.Foreground(m.theme.Colors.Error).Bold(true).Render("Error"),
		m.theme.Styles.Base.Foreground(m.theme.Colors.Foreground).Render(m.err.Error()),
	)
	footer := renderKeybindings(ErrorFetchingSprintKeybindings, m.theme)

	headerHeight := lipgloss.Height(header)
	footerHeight := lipgloss.Height(footer)

	availableHeight := m.height - headerHeight - footerHeight

	body := lipgloss.NewStyle().
		Height(availableHeight).
		Render("")

	return lipgloss.JoinVertical(
		lipgloss.Top,
		header,
		body,
		footer,
	)
}

func footer(m model) string {
	if m.taskDetailsModel.Visible() {
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
