package infrastructure

import (
	"fmt"
	"io/ioutil"
)

// FileWriter provides functionality to write data to a file.
type FileWriter struct{}

// NewFileWriter creates a new instance of FileWriter.
func NewFileWriter() *FileWriter {
	return &FileWriter{}
}

// WriteFile writes the data to the specified file.
func (w *FileWriter) WriteFile(filename string, data []byte) error {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}
	return nil
}
