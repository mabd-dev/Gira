package sprint

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/mabd-dev/gira/internal/theme"
	"github.com/mabd-dev/gira/models"
)

var headerBoxStyle = lipgloss.NewStyle().
	Border(lipgloss.Border{
		Bottom:      "=",
		BottomLeft:  "=",
		BottomRight: "=",
	})

func (m Model) View() string {

	if m.loading {
		loadingStyle := m.theme.Styles.Base.Foreground(m.theme.Colors.Info)
		return loadingStyle.Render("Loading active sprint tasks...")
	}

	if m.err != nil {
		errorStyle := m.theme.Styles.Base.Foreground(m.theme.Colors.Error)
		return errorStyle.Render(fmt.Sprintf("Failed to load projects: %s", m.err.Error()))
	}

	header := header(m.sprint, m.theme)
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

func footer(m Model) string {
	return ""

	// TODO: fix this later
	// if m.taskDetailsModel.Visible() {
	// 	return renderKeybindings(taskdetails.Keybindings, m.theme)
	// } else {
	// 	return renderKeybindings(tasksboard.Keybindings, m.theme)
	// }
}

func developersTabs(m Model) string {
	devs := []string{}
	for i, developer := range m.sprint.Developers {
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
