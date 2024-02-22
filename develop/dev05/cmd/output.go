package cmd

import (
	"fmt"
	"os"
)

// printLines выводит в stdout строки, помеченные в matching.
// withNumbers добавляет номера строк, invert выводит непомеченные строки.
func printLines(lines []string, matching map[int]bool, withNumbers bool, invert bool) {
	for i, line := range lines {
		if withNumbers {
			line = fmt.Sprintf("%d %s", i+1, line)
		}

		if !matching[i] && invert {
			fmt.Fprintln(os.Stdout, line)
			continue
		}

		if matching[i] && !invert {
			fmt.Fprintln(os.Stdout, line)
			continue
		}
	}
}

// printLines выводит в stdout количество строк, помеченных в matching.
//
// Если invert равна true, непомеченных в matching.
func printLineCount(matching map[int]bool, invert bool, linesCount int) {
	var cnt int
	for i := range matching {
		if matching[i] {
			cnt++
		}
	}

	if invert {
		cnt = linesCount - cnt
	}
	fmt.Fprintln(os.Stdout, cnt)
}

// writeError выводит err в stderr.
func writeError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
