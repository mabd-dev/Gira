package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type RealClient struct {
	baseURL    string
	cloudPath  string
	agilePath  string
	username   string
	token      string
	httpClient *http.Client
}

func (c *RealClient) Do(
	method string,
	path string,
	body io.Reader,
) (*http.Response, error) {
	url := c.baseURL + path

	req, err := http.NewRequest(
		method,
		url,
		body,
	)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.username, c.token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("Something went wrong, statusCode=%s, error=%s", res.Status, string(body))
	}

	return res, nil
}

func (c RealClient) GetProjects() (ProjectResponse, error) {
	url := c.cloudPath + "project/search"
	resp, err := c.Do(http.MethodGet, url, nil)
	if err != nil {
		return ProjectResponse{}, err
	}

	getProjectsResponse, err := parseResponse[ProjectResponse](resp)
	if err != nil {
		return ProjectResponse{}, err
	}

	return *getProjectsResponse, nil
}

func (c RealClient) GetBoards(projectID string) (BoardsResponse, error) {
	params := url.Values{}
	params.Add("projectKeyOrId", projectID)

	fullUrl := c.agilePath + "board" + "?" + params.Encode()

	resp, err := c.Do(http.MethodGet, fullUrl, nil)
	if err != nil {
		return BoardsResponse{}, err
	}

	getBoardsResponse, err := parseResponse[BoardsResponse](resp)
	if err != nil {
		return BoardsResponse{}, err
	}

	return *getBoardsResponse, nil
}

func (c RealClient) GetSprints(boardID string) (SprintsResponse, error) {
	fullUrl := c.agilePath + fmt.Sprintf("board/%s/sprint", boardID)

	resp, err := c.Do(http.MethodGet, fullUrl, nil)
	if err != nil {
		return SprintsResponse{}, err
	}

	getSprintsResponse, err := parseResponse[SprintsResponse](resp)
	if err != nil {
		return SprintsResponse{}, err
	}

	return *getSprintsResponse, nil
}

func (c RealClient) GetActiveSprint(boardID string) (Sprint, error) {
	params := url.Values{}
	params.Add("state", "active")

	fullUrl := c.agilePath + fmt.Sprintf("board/%s/sprint", boardID) + "?" + params.Encode()

	resp, err := c.Do(http.MethodGet, fullUrl, nil)
	if err != nil {
		return Sprint{}, err
	}

	getSprintsResponse, err := parseResponse[SprintsResponse](resp)
	if err != nil {
		return Sprint{}, err
	}

	if len(getSprintsResponse.Sprints) == 0 {
		return Sprint{}, fmt.Errorf("No active sprint\n")
	}

	if len(getSprintsResponse.Sprints) > 1 {
		return Sprint{}, fmt.Errorf("More than one active sprint!\n")
	}

	return getSprintsResponse.Sprints[0], nil

}

func (c RealClient) GetSprintIssues(sprintID int) (SprintIssuesResponse, error) {
	fullUrl := c.agilePath + fmt.Sprintf("sprint/%d/issue", sprintID)

	resp, err := c.Do(http.MethodGet, fullUrl, nil)
	if err != nil {
		return SprintIssuesResponse{}, err
	}

	getSprintIssuesResponse, err := parseResponse[SprintIssuesResponse](resp)
	if err != nil {
		return SprintIssuesResponse{}, err
	}

	return *getSprintIssuesResponse, nil
}

func parseResponse[T any](
	response *http.Response,
) (*T, error) {
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return new(T), err
	}

	var parsedResponse T
	if err := json.Unmarshal(responseBody, &parsedResponse); err != nil {
		return new(T), err
	}

	return &parsedResponse, nil
}
