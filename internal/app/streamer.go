package app

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"live-radio-streamer/internal/domain"
	"live-radio-streamer/internal/infrastructure"
)

type Streamer struct {
	streamPath  string
	bufferSize  int
	writer      *infrastructure.BufferWriter
}

func NewStreamer() *Streamer {
	return &Streamer{
		streamPath:  "/",
		bufferSize:  4096,
		writer:      infrastructure.NewBufferWriter(),
	}
}

func (s *Streamer) streamReader(response *http.Response) <-chan []byte {
	ch := make(chan []byte) // Create a channel to receive byte data
	go func() {
		defer close(ch) // Ensure the channel is closed when the goroutine finishes
		startTime := time.Now() // Get the current time as the start time
		for {
			data := make([]byte, s.bufferSize) // Create a byte slice to read the stream data into
			n, err := response.Body.Read(data) // Read data from the response body
			if err != nil {
				if err != io.EOF {
					fmt.Println("Error reading stream:", err)
				}
				break
			}
			ch <- data[:n] // Send the read data to the channel
			if time.Since(startTime).Seconds() >= 20 {
				break
			}
		}
	}()
	return ch // Return the channel to the caller
}


func (s *Streamer) RunStream(radioStreamURL string) (string, error) {
	// Define the directory to store the streamed data
	directory := filepath.Join("internal", "radio_stream_data")

	// Create the directory if it doesn't exist
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Remove "http://" or "https://" from the radio stream URL
	reURL := regexp.MustCompile(`https?://(www\.)?`)
	strippedURL := reURL.ReplaceAllString(radioStreamURL, "")

	// Remove port number from the URL
	strippedURL = regexp.MustCompile(`:\d+`).ReplaceAllString(strippedURL, "")

	// Remove anything after ".com"
	_ = regexp.MustCompile(`\.com.*`).ReplaceAllString(strippedURL, ".com")

	// Generate a unique stream ID based on the current timestamp
	streamID := fmt.Sprintf("%x", time.Now().UnixNano())

	// Make an HTTP GET request to the radio stream URL
	response, err := domain.Get(radioStreamURL)
	if err != nil {
		return "", err
	}

	// Initialize an empty buffer to store the streamed data
	buffer := make([]byte, 0)

	// Iterate over the streamed data and append it to the buffer
	for data := range s.streamReader(response) {
		buffer = append(buffer, data...)
	}

	// Generate the filename with the stream ID
	filename := filepath.Join(directory, streamID+".wav")

	// Write the buffer to a file
	err = s.writer.WriteBuffer(filename, buffer)
	if err != nil {
		return "", err
	}

	// Return the stream ID as a success
	return streamID, nil
}

