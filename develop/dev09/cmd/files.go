package cmd

import (
	"io"
	"os"
)

// safePage записывает body в файл fileName.
func safePage(filePath string, body io.Reader) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, body)
	if err != nil {
		return err
	}

	return nil
}
