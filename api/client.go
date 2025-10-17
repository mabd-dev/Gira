package api

import (
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

func (c *Client) Do(method string, path string, body io.Reader) (*http.Response, error) {
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

	return c.httpClient.Do(req)
}
