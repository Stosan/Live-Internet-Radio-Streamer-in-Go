package main

import (
	"fmt"

	"live-radio-streamer/internal/app"
)

func main() {
	streamer := app.NewStreamer()
	Radio_STREAM_URL := "https://example.com/stream"
	streamID, err := streamer.RunStream(Radio_STREAM_URL)
	if err != nil {
		fmt.Println("Error running stream:", err)
		return
	}
	fmt.Println("Stream ID:", streamID)
}
