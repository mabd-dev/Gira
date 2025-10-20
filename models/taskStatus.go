package models

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/mabd-dev/gira/internal/theme"
)

type TaskStatus string

var (
	TaskStatusTodo       TaskStatus = "Todo"
	TaskStatusInProgress TaskStatus = "In Progress"
	TaskStatusInReview   TaskStatus = "In Review"
	TaskStatusStaging    TaskStatus = "Staging"
	TaskStatusDone       TaskStatus = "Done"
)

var TaskStatusInOrder = []TaskStatus{
	TaskStatusTodo, TaskStatusInProgress, TaskStatusInReview, TaskStatusDone,
}

func getTaskStatusFrom(s string) (TaskStatus, error) {
	switch s {
	case "To Do":
		return TaskStatusTodo, nil
	case "In Progress":
		return TaskStatusInProgress, nil
	case "In Review":
		return TaskStatusInReview, nil
	case "Staging":
		return TaskStatusStaging, nil
	case "Done":
		return TaskStatusDone, nil
	}

	return TaskStatusDone, fmt.Errorf("TaskStatus=%s could not be found!", s)
}

func (t TaskStatus) GetIcon() string {
	switch t {
	case TaskStatusTodo:
		return "○"
	case TaskStatusInProgress:
		return "◐"
	case TaskStatusInReview:
		return "◎"
	case TaskStatusStaging:
		return ""
	case TaskStatusDone:
		return "●"
	}
	return ""
}

func (ts TaskStatus) GetStyle(theme theme.Theme) lipgloss.Style {
	switch ts {
	case TaskStatusTodo:
		return theme.Styles.Base.Bold(true).Foreground(theme.Colors.Muted)
	case TaskStatusInProgress:
		return theme.Styles.Base.Bold(true).Foreground(theme.Colors.Info)
	case TaskStatusInReview:
		return theme.Styles.Base.Bold(true).Foreground(theme.Colors.Warning)
	case TaskStatusDone:
		return theme.Styles.Base.Bold(true).Foreground(theme.Colors.Success)
	}
	return theme.Styles.Base
}
