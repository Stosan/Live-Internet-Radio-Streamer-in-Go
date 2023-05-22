package infrastructure

import (
	"os"
)

// FileWriter provides functionality to write data to a file.
type BufferWriter struct{}

// NewBufferWriter creates a new instance of BufferWriter.
func NewBufferWriter() *BufferWriter {
	return &BufferWriter{}
}

// WriteBuffer writes the data to the specified file.
func (w *BufferWriter) WriteBuffer(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}