package cmd

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

func readLines(file *os.File) ([]string, error) {
	cnt, err := countLines(file)
	if err != nil {
		return nil, err
	}
	file.Seek(0, io.SeekStart)

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0, cnt+1)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func countLines(file *os.File) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}