package api

import (
	"fmt"
	"net/http"
	"net/url"
)

func (c *Client) GetSprints(boardID string) (SprintsResponse, error) {
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

func (c *Client) GetActiveSprint(boardID string) (Sprint, error) {
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

	if len(getSprintsResponse.Sprints) > 1 {
		return Sprint{}, fmt.Errorf("More than one active sprint!\n")
	}

	return getSprintsResponse.Sprints[0], nil

}
