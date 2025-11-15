package api

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// setupTestMockFiles creates mock JSON files in testdata for testing
func setupTestMockFiles(t *testing.T) func() {
	// Save original working directory
	originalWd, err := os.Getwd()
	require.NoError(t, err)

	// Change to parent directory so MockClient can find samples/
	err = os.Chdir("..")
	require.NoError(t, err)

	// Return cleanup function
	return func() {
		os.Chdir(originalWd)
	}
}

func TestNewClient(t *testing.T) {
	username := "test@example.com"
	token := "test-token"
	domain := "test-domain"

	client, err := NewClient(username, token, domain)
	require.NoError(t, err)
	require.NotNil(t, client)

	// Verify the client was set
	assert.NotNil(t, GetClient())
}

func TestNewMockClient(t *testing.T) {
	client, err := NewMockClient()
	require.NoError(t, err)
	require.NotNil(t, client)

	// Verify the client was set
	mockClient := GetClient()
	assert.NotNil(t, mockClient)
	_, ok := mockClient.(MockClient)
	assert.True(t, ok, "Expected client to be MockClient type")
}

func TestMockClient_GetProjects(t *testing.T) {
	cleanup := setupTestMockFiles(t)
	defer cleanup()

	mockClient := MockClient{}

	projects, err := mockClient.GetProjects()
	require.NoError(t, err)
	assert.True(t, projects.IsLast)
	assert.Len(t, projects.Projects, 2)
	assert.Equal(t, "10001", projects.Projects[0].ID)
	assert.Equal(t, "Test Project", projects.Projects[0].Name)
	assert.Equal(t, "software", projects.Projects[0].ProjectTypeKey)
}

func TestMockClient_GetBoards(t *testing.T) {
	cleanup := setupTestMockFiles(t)
	defer cleanup()

	mockClient := MockClient{}

	boards, err := mockClient.GetBoards("10001")
	require.NoError(t, err)
	assert.True(t, boards.IsLast)
	assert.Len(t, boards.Boards, 2)
	assert.Equal(t, 1, boards.Boards[0].ID)
	assert.Equal(t, "Test Board", boards.Boards[0].Name)
	assert.False(t, boards.Boards[0].IsPrivate)
	assert.Equal(t, 2, boards.Boards[1].ID)
	assert.Equal(t, "Private Board", boards.Boards[1].Name)
	assert.True(t, boards.Boards[1].IsPrivate)
}

func TestMockClient_GetSprints(t *testing.T) {
	cleanup := setupTestMockFiles(t)
	defer cleanup()

	mockClient := MockClient{}

	sprints, err := mockClient.GetSprints("1")
	require.NoError(t, err)
	assert.True(t, sprints.IsLast)
	assert.Len(t, sprints.Sprints, 2)
	assert.Equal(t, 1, sprints.Sprints[0].ID)
	assert.Equal(t, "Sprint 1", sprints.Sprints[0].Name)
	assert.Equal(t, "Complete user authentication", sprints.Sprints[0].Goal)
}

func TestMockClient_GetActiveSprint(t *testing.T) {
	t.Run("Returns first sprint as active", func(t *testing.T) {
		cleanup := setupTestMockFiles(t)
		defer cleanup()

		mockClient := MockClient{}

		sprint, err := mockClient.GetActiveSprint("1")
		require.NoError(t, err)
		assert.Equal(t, 1, sprint.ID)
		assert.Equal(t, "Sprint 1", sprint.Name)
		assert.Equal(t, "Complete user authentication", sprint.Goal)
	})
}

func TestMockClient_GetSprintIssues(t *testing.T) {
	cleanup := setupTestMockFiles(t)
	defer cleanup()

	mockClient := MockClient{}

	issues, err := mockClient.GetSprintIssues(1)
	require.NoError(t, err)
	assert.Len(t, issues.Issues, 2)

	// Verify first issue
	assert.Equal(t, "ISSUE-1", issues.Issues[0].ID)
	assert.Equal(t, "John Doe", issues.Issues[0].Fields.Assignee.Name)
	assert.Equal(t, "To Do", issues.Issues[0].Fields.Status.Name)
	assert.Equal(t, "Implement login feature", issues.Issues[0].Fields.Summary)
	assert.Equal(t, float32(5.0), issues.Issues[0].Fields.StoryPoints)
	assert.Len(t, issues.Issues[0].Fields.Components, 2)
	assert.Equal(t, "Backend", issues.Issues[0].Fields.Components[0].Name)
	assert.Len(t, issues.Issues[0].Fields.FixVersions, 1)
	assert.Equal(t, "v1.0.0", issues.Issues[0].Fields.FixVersions[0].Name)

	// Verify second issue
	assert.Equal(t, "ISSUE-2", issues.Issues[1].ID)
	assert.Equal(t, "Jane Smith", issues.Issues[1].Fields.Assignee.Name)
	assert.Equal(t, "In Progress", issues.Issues[1].Fields.Status.Name)
	assert.Equal(t, "Setup database", issues.Issues[1].Fields.Summary)
	assert.Equal(t, float32(3.0), issues.Issues[1].Fields.StoryPoints)
}

func TestGetClient(t *testing.T) {
	// Set up a mock client
	_, err := NewMockClient()
	require.NoError(t, err)

	// Test GetClient returns the set client
	client := GetClient()
	assert.NotNil(t, client)
}
