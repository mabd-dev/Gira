package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	baseUrl    string
	cloudPath  string
	agilePath  string
	username   string
	token      string
	httpClient *http.Client
}

func NewClient(
	username string,
	token string,
	domain string,
) (*Client, error) {
	client := Client{
		baseUrl:    fmt.Sprintf("https://%s.atlassian.net/", domain),
		cloudPath:  "/rest/api/3/",
		agilePath:  "/rest/agile/1.0/",
		username:   username,
		token:      token,
		httpClient: &http.Client{},
	}

	return &client, nil
}

func (c *Client) Do(
	method string,
	path string,
	body io.Reader,
) (*http.Response, error) {
	url := c.baseUrl + path

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

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Something went wrong, statusCode=%s", res.Status)
	}

	return res, nil
}

func parseResponse[T any](
	response *http.Response,
) (*T, error) {

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
