package api

import (
	"net/http"
	"net/url"
)

func (c *Client) GetBoards(projectID string) (BoardsResponse, error) {
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
