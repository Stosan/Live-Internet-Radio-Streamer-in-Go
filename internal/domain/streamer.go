package domain

import (
	c "live-radio-streamer/pkg/http"
	"net/http"
)



// Get makes an HTTP GET request to the specified URL and returns the response.
func Get(url string) (*http.Response, error) {
	c:=c.NewClient()
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

