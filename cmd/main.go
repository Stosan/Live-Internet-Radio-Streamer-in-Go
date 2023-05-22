package main

import (
	"fmt"

	"live-radio-streamer/internal/app"
)

func main() {
	// Create a new instance of Streamer
	streamer := app.NewStreamer()

	// Specify the radio stream URL
	Radio_STREAM_URL := "http://beatfmlagos.atunwadigital.streamguys1.com/beatfmlagos"

	// Run the stream and capture the stream ID
	streamID, err := streamer.RunStream(Radio_STREAM_URL)
	if err != nil {
		fmt.Println("Error running stream:", err)
		return
	}

	// Print the generated stream ID
	fmt.Println("Stream ID:", streamID)
}


