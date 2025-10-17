package api

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	baseUrl    string
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
		baseUrl:    fmt.Sprintf("https://%s.atlassian.net/rest/api/3/", domain),
		username:   username,
		token:      token,
		httpClient: &http.Client{},
	}

	return &client, nil
}

func (c *Client) GetProjects() error {
	url := c.baseUrl + "project/search"

	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.username, c.token)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	// Ensure the response body is closed after reading
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get data, status=%d", res.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Print the response body as a string
	fmt.Println(string(body))

	return nil
}
