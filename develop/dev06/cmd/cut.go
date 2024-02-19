package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func runCut(cmd *cobra.Command, args []string) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		cuttedLine, err := cutLine(line, delimiter, fields, separated)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Fprint(os.Stdout, cuttedLine)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed reading data: %s", err)
		os.Exit(1)
	}
}

func cutLine(line string, delimiter string, fields string, separated bool) (string, error) {
	if len([]rune(delimiter)) != 1 {
		return "", errors.New("bad delimiter")
	}

	if separated {
		if !strings.Contains(line, string(delimiter)) {
			return "", nil // не печатать линию
		}
	}

	indexes, err := getFields(fields)
	if err != nil {
		return "", errors.New("bad fields")
	}

	if line == "" {
		return "\n", nil
	}

	columns := strings.Split(line, delimiter)
	selectedColums := make([]string, 0, len(indexes))

	for _, index := range indexes {
		var column string

		switch {
		case index > len(columns):
			column = columns[0]
		default:
			column = columns[index-1]
		}

		selectedColums = append(selectedColums, column)
	}

	return strings.Join(selectedColums, delimiter) + "\n", nil
}

func getFields(fields string) ([]int, error) {
	digits := strings.Split(fields, ",")
	indexes := make([]int, 0, len(digits))

	for _, d := range digits {
		index, err := strconv.Atoi(d)
		if err != nil {
			return nil, err
		}
		indexes = append(indexes, index)
	}

	return indexes, nil
}
