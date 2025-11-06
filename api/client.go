package api

import (
	"fmt"
	"net/http"
)

type Client interface {
	GetProjects() (ProjectResponse, error)
	GetBoards(projectID string) (BoardsResponse, error)
	GetSprints(boardID string) (SprintsResponse, error)
	GetActiveSprint(boardID string) (Sprint, error)
	GetSprintIssues(sprintID int) (SprintIssuesResponse, error)
}

var client Client

func NewClient(
	username string,
	token string,
	domain string,
) (*Client, error) {
	client = RealClient{
		baseURL:    fmt.Sprintf("https://%s.atlassian.net/", domain),
		cloudPath:  "/rest/api/3/",
		agilePath:  "/rest/agile/1.0/",
		username:   username,
		token:      token,
		httpClient: &http.Client{},
	}

	return &client, nil
}

func NewMockClient() (*Client, error) {
	client = MockClient{}
	return &client, nil
}

func GetClient() Client {
	return client
}
