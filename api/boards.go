package api

import (
	"encoding/json"
	"io"
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

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return BoardsResponse{}, err
	}

	var boardsResponse BoardsResponse
	if err := json.Unmarshal(body, &boardsResponse); err != nil {
		return BoardsResponse{}, err
	}

	return boardsResponse, nil
}
