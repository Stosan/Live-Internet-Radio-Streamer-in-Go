package app

import (
	"fmt"
	"io"
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
	writer      *infrastructure.FileWriter
}

func NewStreamer() *Streamer {
	return &Streamer{
		streamPath:  "/",
		bufferSize:  4096,
		writer:      infrastructure.NewFileWriter(),
	}
}

func (s *Streamer) streamReader(response domain.Response) <-chan []byte {
	ch := make(chan []byte)
	go func() {
		defer close(ch)
		startTime := time.Now()
		for {
			data := make([]byte, s.bufferSize)
			n, err := response.Body().Read(data)
			if err != nil {
				if err != io.EOF {
					fmt.Println("Error reading stream:", err)
				}
				break
			}
			ch <- data[:n]
			if time.Since(startTime).Seconds() >= 1200 {
				break
			}
		}
	}()
	return ch
}

func (s *Streamer) RunStream(radioStreamURL string) (string, error) {
	directory := filepath.Join("src", "internal", "inference", "radio_stream_data")
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return "", err
	}
	reURL := regexp.MustCompile(`https?://(www\.)?`)
	strippedURL := reURL.ReplaceAllString(radioStreamURL, "")
	strippedURL = regexp.MustCompile(`:\d+`).ReplaceAllString(strippedURL, "")
	urlHostname := regexp.MustCompile(`\.com.*`).ReplaceAllString(strippedURL, ".com")

	streamID := fmt.Sprintf("%x", time.Now().UnixNano())
	response, err := domain.Get(radioStreamURL)
	if err != nil {
		return "", err
	}
	defer response.Close()

	buffer := make([]byte, 0)
	for data := range s.streamReader(response) {
		buffer = append(buffer, data...)
	}
	filename := filepath.Join(directory, streamID+".wav")
	err = s.writer.WriteFile(filename, buffer)
	if err != nil {
		return "", err
	}
	return streamID, nil
}
