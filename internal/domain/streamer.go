package domain

import (
	"io"
	"live-radio-streamer/pkg/http"
)

// Response represents the response object from an HTTP request.
type Response interface {
	Body() io.Reader
	Close() error
}

// Get makes an HTTP GET request to the specified URL and returns the response.
func Get(url string) (Response, error) {
	client := http.NewClient()
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
