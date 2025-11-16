package models

import (
	"testing"

	"github.com/mabd-dev/gira/internal/theme"
	"github.com/stretchr/testify/assert"
)

func TestGetTaskStatusFrom_ValidStatuses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected TaskStatus
		wantErr  bool
	}{
		{
			name:     "To Do status",
			input:    "To Do",
			expected: TaskStatusTodo,
			wantErr:  false,
		},
		{
			name:     "In Progress status",
			input:    "In Progress",
			expected: TaskStatusInProgress,
			wantErr:  false,
		},
		{
			name:     "In Review status",
			input:    "In Review",
			expected: TaskStatusInReview,
			wantErr:  false,
		},
		{
			name:     "Staging status",
			input:    "Staging",
			expected: TaskStatusStaging,
			wantErr:  false,
		},
		{
			name:     "Done status",
			input:    "Done",
			expected: TaskStatusDone,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getTaskStatusFrom(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestGetTaskStatusFrom_InvalidStatus(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "Unknown status",
			input: "Unknown Status",
		},
		{
			name:  "Empty string",
			input: "",
		},
		{
			name:  "Random string",
			input: "Not A Real Status",
		},
		{
			name:  "Case sensitive - lowercase",
			input: "to do",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getTaskStatusFrom(tt.input)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), "could not be found")
			// Should default to Done when not found
			assert.Equal(t, TaskStatusDone, result)
		})
	}
}

func TestTaskStatus_GetIcon(t *testing.T) {
	tests := []struct {
		name     string
		status   TaskStatus
		expected string
	}{
		{
			name:     "Todo icon",
			status:   TaskStatusTodo,
			expected: "○",
		},
		{
			name:     "In Progress icon",
			status:   TaskStatusInProgress,
			expected: "◐",
		},
		{
			name:     "In Review icon",
			status:   TaskStatusInReview,
			expected: "◎",
		},
		{
			name:     "Staging icon",
			status:   TaskStatusStaging,
			expected: "",
		},
		{
			name:     "Done icon",
			status:   TaskStatusDone,
			expected: "●",
		},
		{
			name:     "Unknown status icon",
			status:   TaskStatus("Unknown"),
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.status.GetIcon()
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestTaskStatus_GetStyle(t *testing.T) {
	// Create a theme for testing
	colors, err := theme.CreateColors("catppuccin-mocha")
	assert.NoError(t, err)
	styles := theme.CreateStyles(colors)
	th := theme.Theme{
		Colors: colors,
		Styles: styles,
	}

	tests := []struct {
		name   string
		status TaskStatus
	}{
		{
			name:   "Todo style",
			status: TaskStatusTodo,
		},
		{
			name:   "In Progress style",
			status: TaskStatusInProgress,
		},
		{
			name:   "In Review style",
			status: TaskStatusInReview,
		},
		{
			name:   "Done style",
			status: TaskStatusDone,
		},
		{
			name:   "Unknown status style",
			status: TaskStatus("Unknown"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Just verify it doesn't panic and returns a style
			style := tt.status.GetStyle(th)
			assert.NotNil(t, style)
		})
	}
}

func TestTaskStatusInOrder(t *testing.T) {
	// Verify the TaskStatusInOrder slice has the expected statuses in order
	expected := []TaskStatus{
		TaskStatusTodo,
		TaskStatusInProgress,
		TaskStatusInReview,
		TaskStatusDone,
	}
	assert.Equal(t, expected, TaskStatusInOrder)
}
