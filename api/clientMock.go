package api

import (
	"encoding/json"
	"fmt"
	"os"
)

type MockClient struct{}

func (c MockClient) GetProjects() (ProjectResponse, error) {
	body, err := readMockFile[ProjectResponse]("samples/mockApiResponses/projects.json")
	return *body, err
}

func (c MockClient) GetBoards(projectID string) (BoardsResponse, error) {
	body, err := readMockFile[BoardsResponse]("samples/mockApiResponses/boards.json")
	return *body, err
}

func (c MockClient) GetSprints(boardID string) (SprintsResponse, error) {
	body, err := readMockFile[SprintsResponse]("samples/mockApiResponses/sprints.json")
	return *body, err
}

func (c MockClient) GetActiveSprint(boardID string) (Sprint, error) {
	body, err := readMockFile[SprintsResponse]("samples/mockApiResponses/sprints.json")
	if err != nil {
		return Sprint{}, err
	}

	if len(body.Sprints) == 0 {
		return Sprint{}, fmt.Errorf("no active sprint found")
	}

	return body.Sprints[0], nil
}

func (c MockClient) GetSprintIssues(sprintID int) (SprintIssuesResponse, error) {
	body, err := readMockFile[SprintIssuesResponse]("samples/mockApiResponses/sprintIssues.json")
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
