package http

import (
	"net/http"
	"time"
)

// Client provides functionality for making HTTP requests.
type Client struct {
	client *http.Client
}

// NewClient creates a new instance of the HTTP client.
func NewClient() *Client {
	return &Client{
		client: &http.Client{
			Timeout: 10 * time.Second, // Set a timeout for requests
		},
	}
}

// Get performs an HTTP GET request to the specified URL and returns the response.
func (c *Client) Get(url string) (*http.Response, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
