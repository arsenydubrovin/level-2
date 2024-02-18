package cmd

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

// loadDictionary возвращает множество строк, прочитанных из file
func LoadDictionary(file *os.File) (map[string]struct{}, error) {
	linesCount, err := countLines(file)
	if err != nil {
		return nil, err
	}

	dict := make(map[string]struct{}, linesCount)

	file.Seek(0, io.SeekStart) // чтобы прочитать из файла повторно
	readLines(file, dict)

	return dict, nil
}

// countLines считывает строки из источника r в мапу m.
func readLines(r io.Reader, m map[string]struct{}) error {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		m[scanner.Text()] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// countLines считает количество строк в источнике.
func countLines(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
