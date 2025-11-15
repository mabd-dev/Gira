package api

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type MockClient struct {
	BasePath string
}

// api/testdata/mockApiResponses/

func (c MockClient) GetProjects() (ProjectResponse, error) {
	path := filepath.Join(c.BasePath, "api/testdata/mockApiResponses/projects.json")
	body, err := readMockFile[ProjectResponse](path)
	return *body, err
}

func (c MockClient) GetBoards(projectID string) (BoardsResponse, error) {
	path := filepath.Join(c.BasePath, "api/testdata/mockApiResponses/boards.json")
	body, err := readMockFile[BoardsResponse](path)
	return *body, err
}

func (c MockClient) GetSprints(boardID string) (SprintsResponse, error) {
	path := filepath.Join(c.BasePath, "api/testdata/mockApiResponses/sprints.json")
	body, err := readMockFile[SprintsResponse](path)
	return *body, err
}

func (c MockClient) GetActiveSprint(boardID string) (Sprint, error) {
	path := filepath.Join(c.BasePath, "api/testdata/mockApiResponses/sprints.json")
	body, err := readMockFile[SprintsResponse](path)
	if err != nil {
		return Sprint{}, err
	}

	if len(body.Sprints) == 0 {
		return Sprint{}, fmt.Errorf("no active sprint found")
	}

	return body.Sprints[0], nil
}

func (c MockClient) GetSprintIssues(sprintID int) (SprintIssuesResponse, error) {
	path := filepath.Join(c.BasePath, "api/testdata/mockApiResponses/sprintIssues.json")
	body, err := readMockFile[SprintIssuesResponse](path)
	return *body, err
}

func readMockFile[T any](filePath string) (*T, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return new(T), err
	}

	var response T
	if err := json.Unmarshal(fileData, &response); err != nil {
		return new(T), err
	}

	return &response, nil
}
