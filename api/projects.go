package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetProjects() (ProjectResponse, error) {
	url := c.cloudPath + "project/search"
	res, err := c.Do(http.MethodGet, url, nil)
	if err != nil {
		return ProjectResponse{}, err
	}

	// Ensure the response body is closed after reading
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ProjectResponse{}, fmt.Errorf("failed to get data, status=%d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ProjectResponse{}, err
	}

	var projects ProjectResponse
	if err := json.Unmarshal(body, &projects); err != nil {
		return ProjectResponse{}, err
	}

	return projects, nil
}
