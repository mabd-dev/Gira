package api

import (
	"net/http"
)

func (c *Client) GetProjects() (ProjectResponse, error) {
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
