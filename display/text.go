package display

import (
	"fmt"
	"strings"
	"time"

	"github.com/mabd-dev/gira/models"
)

// ANSI color codes
const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
	Dim   = "\033[2m"

	// Foreground colors
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	// Background colors
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
)

func PrintSprint(sprint models.Sprint) {
	// Print sprint header
	printSprintHeader(sprint)

	// Print each developer's tasks
	for i, dev := range sprint.Developers {
		printDeveloper(dev)

		// Add spacing between developers (except for the last one)
		if i < len(sprint.Developers)-1 {
			fmt.Println()
		}
	}

	// Print footer
	printFooter()
}

func printSprintHeader(sprint models.Sprint) {
	width := 80
	line := strings.Repeat("━", width)

	fmt.Println(Bold + Cyan + line + Reset)

	// Sprint name and remaining days
	remainingDays := getRemainingDays(sprint.EndDate)
	daysColor := Green
	if remainingDays <= 3 {
		daysColor = Red
	} else if remainingDays <= 7 {
		daysColor = Yellow
	}

	fmt.Printf("%s%s  🏃 SPRINT: %s%s\n", Bold, White, sprint.Name, Reset)
	fmt.Printf("%s%s  ⏰ Days Remaining: %s%s%d%s\n",
		Bold, White, daysColor, Bold, remainingDays, Reset)

	// Sprint goal
	goal := formatSprintGoal(sprint.Goal)
	fmt.Printf("%s%s  🎯 Goal: %s%s\n", Bold, White, Yellow+goal, Reset)

	fmt.Println(Bold + Cyan + line + Reset)
	fmt.Println()
}

func printDeveloper(dev models.Developer) {
	// Developer name header
	fmt.Printf("%s%s┌─ %s%s\n", Dim, White, dev.Name, Reset)

	hasAnyTasks := false
	for _, status := range models.TasksInOrder {
		tasks := dev.TasksByStatus[status]
		if len(tasks) > 0 {
			hasAnyTasks = true
			break
		}
	}

	if !hasAnyTasks {
		fmt.Printf("%s%s│  %s(No tasks assigned)%s\n", Bold, Blue, Dim, Reset)
		fmt.Printf("%s%s└─%s\n", Bold, Blue, Reset)
		return
	}

	// Print tasks by status
	statusCount := 0
	totalStatuses := 0
	for _, status := range models.TasksInOrder {
		if len(dev.TasksByStatus[status]) > 0 {
			totalStatuses++
		}
	}

	for _, status := range models.TasksInOrder {
		tasks := dev.TasksByStatus[status]
		if len(tasks) == 0 {
			continue
		}

		statusCount++
		isLast := statusCount == totalStatuses

		printTaskStatus(status, tasks, isLast)
	}

	fmt.Printf("%s%s└─%s\n", Bold, Blue, Reset)
}

func printTaskStatus(status models.TaskStatus, tasks []models.DeveloperTask, isLast bool) {
	// Get color based on status
	statusColor := getStatusColor(status)
	connector := "├─"

	// Calculate total story points
	totalPoints := 0
	for _, task := range tasks {
		totalPoints += task.StoryPoints
	}

	// Print status header
	fmt.Printf("%s%s│%s\n", Bold, Blue, Reset)
	fmt.Printf("%s%s%s %s%s %s (%d tasks, %d pts)%s\n",
		Bold, Blue, connector, statusColor, status, status.GetIcon(),
		len(tasks), totalPoints, Reset)

	// Print each task
	for i, task := range tasks {
		pointsColor := getPointsColor(task.StoryPoints)

		taskPrefix := "│  "
		if isLast && i == len(tasks)-1 {
			taskPrefix = "│  "
		}

		// Truncate long summaries
		summary := task.Summary
		// maxLen := 65
		// if len(summary) > maxLen {
		// 	summary = summary[:maxLen-3] + "..."
		// }

		fmt.Printf("%s%s%s   %s%d.%s %s[%s%d%s] %s%s\n",
			Bold, Blue, taskPrefix,
			Dim, i+1, Reset,
			pointsColor, Bold, task.StoryPoints, Reset,
			White, summary)
	}
}

func printFooter() {
	width := 80
	line := strings.Repeat("━", width)
	fmt.Println(Bold + Cyan + line + Reset)
}

func getStatusColor(status models.TaskStatus) string {
	switch status {
	case models.TaskStatusTodo:
		return White
	case models.TaskStatusInProgress:
		return Yellow
	case models.TaskStatusInReview:
		return Magenta
	case models.TaskStatusStaging:
		return Cyan
	case models.TaskStatusDone:
		return Green
	default:
		return White
	}
}

func getPointsColor(points int) string {
	if points >= 8 {
		return Red
	} else if points >= 5 {
		return Yellow
	} else if points >= 3 {
		return Cyan
	}
	return Green
}

func getRemainingDays(endDate string) int {
	// Parse the end date - try common formats
	var sprintEnd time.Time
	var err error

	// Try ISO 8601 format first (e.g., "2025-10-24T17:00:00.000Z")
	sprintEnd, err = time.Parse(time.RFC3339, endDate)
	if err != nil {
		// Try date-only format (e.g., "2025-10-24")
		sprintEnd, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			// If parsing fails, return 0
			return 0
		}
	}

	// Get current time
	now := time.Now()

	// Calculate the difference in days
	duration := sprintEnd.Sub(now)
	days := int(duration.Hours() / 24)

	// Return 0 if negative (sprint already ended)
	if days < 0 {
		return 0
	}

	return days
}

func formatSprintGoal(goal string) string {
	if len(strings.TrimSpace(goal)) == 0 {
		return "Nothing !"
	}
	return goal
}
