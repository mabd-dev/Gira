package models

import (
	"testing"

	"github.com/mabd-dev/gira/api"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatProject(t *testing.T) {
	t.Run("Format single project", func(t *testing.T) {
		apiResponse := api.ProjectResponse{
			IsLast: true,
			Projects: []api.Project{
				{
					ID:             "10001",
					Name:           "Test Project",
					ProjectTypeKey: "software",
				},
			},
		}

		result, err := FormatProject(apiResponse)
		require.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "10001", result[0].ID)
		assert.Equal(t, "Test Project", result[0].Name)
		assert.Equal(t, "software", result[0].ProjectTypeKey)
	})

	t.Run("Format multiple projects", func(t *testing.T) {
		apiResponse := api.ProjectResponse{
			IsLast: false,
			Projects: []api.Project{
				{
					ID:             "10001",
					Name:           "Project One",
					ProjectTypeKey: "software",
				},
				{
					ID:             "10002",
					Name:           "Project Two",
					ProjectTypeKey: "business",
				},
			},
		}

		result, err := FormatProject(apiResponse)
		require.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, "10001", result[0].ID)
		assert.Equal(t, "Project One", result[0].Name)
		assert.Equal(t, "10002", result[1].ID)
		assert.Equal(t, "Project Two", result[1].Name)
	})

	t.Run("Format empty project list", func(t *testing.T) {
		apiResponse := api.ProjectResponse{
			IsLast:   true,
			Projects: []api.Project{},
		}

		result, err := FormatProject(apiResponse)
		require.NoError(t, err)
		assert.Len(t, result, 0)
	})
}

func TestFormatBoards(t *testing.T) {
	t.Run("Format single board", func(t *testing.T) {
		apiResponse := api.BoardsResponse{
			IsLast: true,
			Boards: []api.Board{
				{
					ID:        1,
					Name:      "Test Board",
					IsPrivate: true,
				},
			},
		}

		result, err := FormatBoards(apiResponse)
		require.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, 1, result[0].ID)
		assert.Equal(t, "Test Board", result[0].Name)
		assert.True(t, result[0].IsPrivate)
	})

	t.Run("Format multiple boards", func(t *testing.T) {
		apiResponse := api.BoardsResponse{
			IsLast: false,
			Boards: []api.Board{
				{
					ID:        1,
					Name:      "Board One",
					IsPrivate: true,
				},
				{
					ID:        2,
					Name:      "Board Two",
					IsPrivate: false,
				},
			},
		}

		result, err := FormatBoards(apiResponse)
		require.NoError(t, err)
		assert.Len(t, result, 2)
		assert.Equal(t, 1, result[0].ID)
		assert.Equal(t, "Board One", result[0].Name)
		assert.True(t, result[0].IsPrivate)
		assert.Equal(t, 2, result[1].ID)
		assert.Equal(t, "Board Two", result[1].Name)
		assert.False(t, result[1].IsPrivate)
	})

	t.Run("Format empty board list", func(t *testing.T) {
		apiResponse := api.BoardsResponse{
			IsLast: true,
			Boards: []api.Board{},
		}

		result, err := FormatBoards(apiResponse)
		require.NoError(t, err)
		assert.Len(t, result, 0)
	})
}

func TestFormatSprint(t *testing.T) {
	t.Run("Format sprint with single developer and task", func(t *testing.T) {
		apiResponse := api.SprintIssuesResponse{
			Issues: []api.Issue{
				{
					ID: "ISSUE-1",
					Fields: api.IssueFields{
						Assignee: api.IssueAssignee{
							AccountID: "123",
							Name:      "John Doe",
						},
						Status: api.IssueStatus{
							Name: "To Do",
						},
						Summary:     "Test task",
						Description: "Test description",
						Sprint: api.Sprint{
							ID:        1,
							Name:      "Sprint 1",
							StartDate: "2024-01-01",
							EndDate:   "2024-01-14",
							Goal:      "Sprint goal",
						},
						StoryPoints: 5.0,
						Components: []api.IssueComponent{
							{ID: "1", Name: "Backend"},
						},
						FixVersions: []api.IssueFixVersion{
							{ID: "1", Name: "v1.0.0", Description: "First release"},
						},
					},
				},
			},
		}

		result, err := FormatSprint(apiResponse)
		require.NoError(t, err)
		assert.Equal(t, "ISSUE-1", result.ID)
		assert.Equal(t, "Sprint 1", result.Name)
		assert.Equal(t, "2024-01-01", result.StartDate)
		assert.Equal(t, "2024-01-14", result.EndDate)
		assert.Equal(t, "Sprint goal", result.Goal)
		assert.Len(t, result.Developers, 1)
		assert.Equal(t, "John Doe", result.Developers[0].Name)
		assert.Len(t, result.Developers[0].TasksByStatus[TaskStatusTodo], 1)
		assert.Equal(t, "ISSUE-1", result.Developers[0].TasksByStatus[TaskStatusTodo][0].ID)
		assert.Equal(t, "Test task", result.Developers[0].TasksByStatus[TaskStatusTodo][0].Summary)
		assert.Equal(t, 5, result.Developers[0].TasksByStatus[TaskStatusTodo][0].StoryPoints)
		assert.Equal(t, []string{"Backend"}, result.Developers[0].TasksByStatus[TaskStatusTodo][0].Components)
		assert.Equal(t, []string{"v1.0.0"}, result.Developers[0].TasksByStatus[TaskStatusTodo][0].FixVersions)
	})

	t.Run("Format sprint with multiple developers", func(t *testing.T) {
		apiResponse := api.SprintIssuesResponse{
			Issues: []api.Issue{
				{
					ID: "ISSUE-1",
					Fields: api.IssueFields{
						Assignee: api.IssueAssignee{Name: "Alice"},
						Status:   api.IssueStatus{Name: "To Do"},
						Summary:  "Task 1",
						Sprint: api.Sprint{
							ID:        1,
							Name:      "Sprint 1",
							StartDate: "2024-01-01",
							EndDate:   "2024-01-14",
							Goal:      "Sprint goal",
						},
						StoryPoints: 3.0,
					},
				},
				{
					ID: "ISSUE-2",
					Fields: api.IssueFields{
						Assignee: api.IssueAssignee{Name: "Bob"},
						Status:   api.IssueStatus{Name: "In Progress"},
						Summary:  "Task 2",
						Sprint: api.Sprint{
							ID:        1,
							Name:      "Sprint 1",
							StartDate: "2024-01-01",
							EndDate:   "2024-01-14",
							Goal:      "Sprint goal",
						},
						StoryPoints: 5.0,
					},
				},
			},
		}

		result, err := FormatSprint(apiResponse)
		require.NoError(t, err)
		assert.Len(t, result.Developers, 2)
		// Developers should be sorted alphabetically
		assert.Equal(t, "Alice", result.Developers[0].Name)
		assert.Equal(t, "Bob", result.Developers[1].Name)
	})

	t.Run("Format sprint with unassigned tasks", func(t *testing.T) {
		apiResponse := api.SprintIssuesResponse{
			Issues: []api.Issue{
				{
					ID: "ISSUE-1",
					Fields: api.IssueFields{
						Assignee: api.IssueAssignee{Name: ""}, // Empty assignee
						Status:   api.IssueStatus{Name: "To Do"},
						Summary:  "Unassigned task",
						Sprint: api.Sprint{
							ID:        1,
							Name:      "Sprint 1",
							StartDate: "2024-01-01",
							EndDate:   "2024-01-14",
						},
						StoryPoints: 2.0,
					},
				},
			},
		}

		result, err := FormatSprint(apiResponse)
		require.NoError(t, err)
		assert.Len(t, result.Developers, 1)
		assert.Equal(t, "Unassigned", result.Developers[0].Name)
	})

	t.Run("Format sprint with same developer multiple tasks", func(t *testing.T) {
		apiResponse := api.SprintIssuesResponse{
			Issues: []api.Issue{
				{
					ID: "ISSUE-1",
					Fields: api.IssueFields{
						Assignee: api.IssueAssignee{Name: "Alice"},
						Status:   api.IssueStatus{Name: "To Do"},
						Summary:  "Task 1",
						Sprint: api.Sprint{
							ID:   1,
							Name: "Sprint 1",
						},
						StoryPoints: 3.0,
					},
				},
				{
					ID: "ISSUE-2",
					Fields: api.IssueFields{
						Assignee: api.IssueAssignee{Name: "Alice"},
						Status:   api.IssueStatus{Name: "In Progress"},
						Summary:  "Task 2",
						Sprint: api.Sprint{
							ID:   1,
							Name: "Sprint 1",
						},
						StoryPoints: 5.0,
					},
				},
				{
					ID: "ISSUE-3",
					Fields: api.IssueFields{
						Assignee: api.IssueAssignee{Name: "Alice"},
						Status:   api.IssueStatus{Name: "To Do"},
						Summary:  "Task 3",
						Sprint: api.Sprint{
							ID:   1,
							Name: "Sprint 1",
						},
						StoryPoints: 2.0,
					},
				},
			},
		}

		result, err := FormatSprint(apiResponse)
		require.NoError(t, err)
		assert.Len(t, result.Developers, 1)
		assert.Equal(t, "Alice", result.Developers[0].Name)
		assert.Len(t, result.Developers[0].TasksByStatus[TaskStatusTodo], 2)
		assert.Len(t, result.Developers[0].TasksByStatus[TaskStatusInProgress], 1)
	})

	t.Run("Format sprint with invalid status", func(t *testing.T) {
		apiResponse := api.SprintIssuesResponse{
			Issues: []api.Issue{
				{
					ID: "ISSUE-1",
					Fields: api.IssueFields{
						Assignee: api.IssueAssignee{Name: "Alice"},
						Status:   api.IssueStatus{Name: "Invalid Status"},
						Summary:  "Task with invalid status",
						Sprint: api.Sprint{
							ID:   1,
							Name: "Sprint 1",
						},
						StoryPoints: 3.0,
					},
				},
			},
		}

		result, err := FormatSprint(apiResponse)
		require.NoError(t, err)
		// Invalid status is skipped but developer entry is still created with empty tasks
		assert.Len(t, result.Developers, 1)
		assert.Equal(t, "Alice", result.Developers[0].Name)
		// TasksByStatus should be empty since the task was skipped
		assert.Equal(t, 0, len(result.Developers[0].TasksByStatus))
	})

	t.Run("Format empty sprint", func(t *testing.T) {
		apiResponse := api.SprintIssuesResponse{
			Issues: []api.Issue{},
		}

		result, err := FormatSprint(apiResponse)
		require.NoError(t, err)
		assert.Equal(t, Sprint{}, result)
	})
}

func TestParseIssueComponents(t *testing.T) {
	t.Run("Parse single component", func(t *testing.T) {
		components := []api.IssueComponent{
			{ID: "1", Name: "Backend"},
		}
		result := parseIssueComponents(components)
		assert.Equal(t, []string{"Backend"}, result)
	})

	t.Run("Parse multiple components", func(t *testing.T) {
		components := []api.IssueComponent{
			{ID: "1", Name: "Backend"},
			{ID: "2", Name: "Frontend"},
			{ID: "3", Name: "Database"},
		}
		result := parseIssueComponents(components)
		assert.Equal(t, []string{"Backend", "Frontend", "Database"}, result)
	})

	t.Run("Parse empty components", func(t *testing.T) {
		components := []api.IssueComponent{}
		result := parseIssueComponents(components)
		assert.Equal(t, []string{}, result)
	})
}

func TestParseIssueFixVersions(t *testing.T) {
	t.Run("Parse single fix version", func(t *testing.T) {
		versions := []api.IssueFixVersion{
			{ID: "1", Name: "v1.0.0", Description: "First release"},
		}
		result := parseIssueFixVersions(versions)
		assert.Equal(t, []string{"v1.0.0"}, result)
	})

	t.Run("Parse multiple fix versions", func(t *testing.T) {
		versions := []api.IssueFixVersion{
			{ID: "1", Name: "v1.0.0", Description: "First release"},
			{ID: "2", Name: "v1.1.0", Description: "Second release"},
		}
		result := parseIssueFixVersions(versions)
		assert.Equal(t, []string{"v1.0.0", "v1.1.0"}, result)
	})

	t.Run("Parse empty fix versions", func(t *testing.T) {
		versions := []api.IssueFixVersion{}
		result := parseIssueFixVersions(versions)
		assert.Equal(t, []string{}, result)
	})
}
