package taskdetails

import (
	"regexp"
	"strings"

	"github.com/mabd-dev/gira/internal/theme"
)

// formatJiraDescription converts Jira wiki markup to styled terminal output
func formatJiraDescription(description string, theme theme.Theme) string {
	if description == "" {
		return ""
	}

	lines := strings.Split(description, "\n")
	var formattedLines []string

	// Styles
	h2Style := theme.Styles.Base.
		Foreground(theme.Colors.Foreground).
		Bold(true).
		Underline(true)

	normalStyle := theme.Styles.Base.
		Foreground(theme.Colors.Foreground)

	listItemStyle := theme.Styles.Base.
		Foreground(theme.Colors.Foreground).
		PaddingLeft(2)

	subListItemStyle := theme.Styles.Base.
		Foreground(theme.Colors.Foreground).
		PaddingLeft(4)

	// Regex patterns for Jira markup
	h2Pattern := regexp.MustCompile(`^h2\.\s*(.+)$`)
	boldPattern := regexp.MustCompile(`\*([^*]+)\*`)
	strikethroughPattern := regexp.MustCompile(`-([^*]+)-`)
	linkPattern := regexp.MustCompile(`\[([^\|]+)\|([^\]]+)\]`)
	numberedListPattern := regexp.MustCompile(`^#\s+(.+)$`)
	bulletListPattern := regexp.MustCompile(`^\*\s+(.+)$`)
	tableRowPattern := regexp.MustCompile(`^\|.+\|$`)

	inTable := false
	var tableRows []string

	for _, line := range lines {
		line = strings.TrimRight(line, " ")

		// Skip empty lines but preserve spacing
		if line == "" {
			// If we were in a table, render it now
			if inTable {
				formattedLines = append(formattedLines, formatTable(tableRows, theme))
				tableRows = nil
				inTable = false
			}
			formattedLines = append(formattedLines, "")
			continue
		}

		// Handle tables
		if tableRowPattern.MatchString(line) {
			inTable = true
			tableRows = append(tableRows, line)
			continue
		} else if inTable {
			// End of table, render it
			formattedLines = append(formattedLines, formatTable(tableRows, theme))
			tableRows = nil
			inTable = false
		}

		// Handle h2 headers
		if h2Pattern.MatchString(line) {
			matches := h2Pattern.FindStringSubmatch(line)
			if len(matches) > 1 {
				headerText := formatInlineMarkup(matches[1], boldPattern, strikethroughPattern, linkPattern, theme)
				formattedLines = append(formattedLines, h2Style.Render(headerText))
				continue
			}
		}

		// Handle numbered lists (# at start)
		if numberedListPattern.MatchString(line) {
			matches := numberedListPattern.FindStringSubmatch(line)
			if len(matches) > 1 {
				content := formatInlineMarkup(matches[1], boldPattern, strikethroughPattern, linkPattern, theme)
				formattedLines = append(formattedLines, listItemStyle.Render("• "+content))
				continue
			}
		}

		// Handle bullet lists (* at start of line)
		if bulletListPattern.MatchString(line) {
			matches := bulletListPattern.FindStringSubmatch(line)
			if len(matches) > 1 {
				content := formatInlineMarkup(matches[1], boldPattern, strikethroughPattern, linkPattern, theme)
				formattedLines = append(formattedLines, subListItemStyle.Render("◦ "+content))
				continue
			}
		}

		// Handle regular text with inline formatting
		formatted := formatInlineMarkup(line, boldPattern, strikethroughPattern, linkPattern, theme)
		formattedLines = append(formattedLines, normalStyle.Render(formatted))
	}

	// Handle table at end of document
	if inTable {
		formattedLines = append(formattedLines, formatTable(tableRows, theme))
	}

	return strings.Join(formattedLines, "\n")
}

// formatInlineMarkup handles inline formatting like *bold* and [text|url]
func formatInlineMarkup(text string, boldPattern *regexp.Regexp, strikethroughPattern *regexp.Regexp, linkPattern *regexp.Regexp, theme theme.Theme) string {
	result := text

	// Handle links first: [text|url] -> "text (url)"
	if matches := linkPattern.FindAllStringSubmatch(text, -1); len(matches) > 0 {
		for _, match := range matches {
			if len(match) > 2 {
				fullMatch := match[0] // [text|url]
				linkText := match[1]  // text
				linkURL := match[2]   // url

				// Style: text in normal color, URL in muted/blue
				linkStyle := theme.Styles.Base.Foreground(theme.Colors.Foreground)
				urlStyle := theme.Styles.Base.Foreground(theme.Colors.Accent).Underline(true)

				styledLink := linkStyle.Render(linkText) + " " + urlStyle.Render("("+linkURL+")")
				result = strings.Replace(result, fullMatch, styledLink, 1)
			}
		}
	}

	// Handle bold text: *text*
	if matches := boldPattern.FindAllStringSubmatch(result, -1); len(matches) > 0 {
		for _, match := range matches {
			if len(match) > 1 {
				fullMatch := match[0] // *text*
				boldText := match[1]  // text
				styledText := theme.Styles.Base.Bold(true).Foreground(theme.Colors.Foreground).Render(boldText)
				result = strings.Replace(result, fullMatch, styledText, 1)
			}
		}
	}

	// Handle strikethrough text: -text-
	if matches := strikethroughPattern.FindAllStringSubmatch(result, -1); len(matches) > 0 {
		for _, match := range matches {
			if len(match) > 1 {
				fullMatch := match[0] // -text-
				boldText := match[1]  // text
				styledText := theme.Styles.Base.Strikethrough(true).Foreground(theme.Colors.Foreground).Render(boldText)
				result = strings.Replace(result, fullMatch, styledText, 1)
			}
		}
	}

	return result
}

// formatTable converts Jira table markup to formatted terminal table
func formatTable(rows []string, theme theme.Theme) string {
	if len(rows) == 0 {
		return ""
	}

	// Parse table rows
	var parsedRows [][]string
	isHeaderRow := false

	for _, row := range rows {
		// Remove leading and trailing pipes
		row = strings.Trim(row, "|")

		// Split by pipe
		cells := strings.Split(row, "|")

		// Trim whitespace from each cell
		for i, cell := range cells {
			cells[i] = strings.TrimSpace(cell)
		}

		// Check if this is a header row (contains ||)
		if strings.Contains(rows[0], "||") && len(parsedRows) == 0 {
			isHeaderRow = true
			// For header rows, split by || instead
			row = rows[0]
			row = strings.Trim(row, "|")
			cells = strings.Split(row, "||")
			for i, cell := range cells {
				cells[i] = strings.TrimSpace(cell)
			}
		}

		parsedRows = append(parsedRows, cells)
		if isHeaderRow {
			isHeaderRow = false
		}
	}

	// Calculate column widths
	colWidths := make([]int, 0)
	for _, row := range parsedRows {
		for i, cell := range row {
			// Ensure colWidths is long enough
			for len(colWidths) <= i {
				colWidths = append(colWidths, 0)
			}
			if len(cell) > colWidths[i] {
				colWidths[i] = len(cell)
			}
		}
	}

	// Styles
	headerStyle := theme.Styles.Base.Foreground(theme.Colors.Foreground).Bold(true)
	cellStyle := theme.Styles.Base.Foreground(theme.Colors.Foreground)
	borderStyle := theme.Styles.Muted

	// Build table
	var tableLines []string

	for rowIdx, row := range parsedRows {
		var cellsFormatted []string

		for colIdx, cell := range row {
			// Pad cell to column width
			width := colWidths[colIdx]
			paddedCell := cell + strings.Repeat(" ", width-len(cell))

			// Apply style based on whether it's header row
			if rowIdx == 0 && strings.Contains(rows[0], "||") {
				cellsFormatted = append(cellsFormatted, headerStyle.Render(paddedCell))
			} else {
				cellsFormatted = append(cellsFormatted, cellStyle.Render(paddedCell))
			}
		}

		// Join cells with border
		border := borderStyle.Render(" │ ")
		tableLine := borderStyle.Render("│ ") + strings.Join(cellsFormatted, border) + borderStyle.Render(" │")
		tableLines = append(tableLines, tableLine)

		// Add separator after header row
		if rowIdx == 0 && strings.Contains(rows[0], "||") {
			var separatorParts []string
			for _, width := range colWidths {
				separatorParts = append(separatorParts, strings.Repeat("─", width))
			}
			separator := borderStyle.Render("├─") + borderStyle.Render(strings.Join(separatorParts, "─┼─")) + borderStyle.Render("─┤")
			tableLines = append(tableLines, separator)
		}
	}

	// Add top border
	var topParts []string
	for _, width := range colWidths {
		topParts = append(topParts, strings.Repeat("─", width))
	}
	topBorder := borderStyle.Render("┌─") + borderStyle.Render(strings.Join(topParts, "─┬─")) + borderStyle.Render("─┐")

	// Add bottom border
	var bottomParts []string
	for _, width := range colWidths {
		bottomParts = append(bottomParts, strings.Repeat("─", width))
	}
	bottomBorder := borderStyle.Render("└─") + borderStyle.Render(strings.Join(bottomParts, "─┴─")) + borderStyle.Render("─┘")

	// Combine everything
	result := []string{topBorder}
	result = append(result, tableLines...)
	result = append(result, bottomBorder)

	return strings.Join(result, "\n")
}
