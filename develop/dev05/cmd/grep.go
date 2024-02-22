package cmd

import (
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

// runGrep основную логику программы
func runGrep(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		writeError(cmd.Usage())
	}

	cfg, err := readConfig(cmd)
	if err != nil {
		writeError(err)
	}

	file, err := os.Open(args[1])
	if err != nil {
		writeError(err)
	}
	defer file.Close()

	lines, err := readLines(file)
	if err != nil {
		writeError(err)
	}

	matchingIndexes, err := getMatching(lines, args[0], cfg.fixed, cfg.ignoreCase)
	if err != nil {
		writeError(err)
	}

	if cfg.count {
		printLineCount(matchingIndexes, cfg.invert, len(lines))
		return
	}

	markIndexesToPrint(matchingIndexes, cfg.after, cfg.before, cfg.context, len(lines))

	printLines(lines, matchingIndexes, cfg.lineNumbers, cfg.invert)
}

// getMatching помечает как true индексы строк, которые соответсвуют pattern
func getMatching(lines []string, pattern string, fixed bool, ignoreCase bool) (map[int]bool, error) {
	if ignoreCase {
		pattern = "(?i)" + pattern
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	matchingIndexes := make(map[int]bool, 0)
	for i, line := range lines {
		if !fixed && re.MatchString(line) || fixed && line == pattern {
			matchingIndexes[i] = true
		}
	}

	return matchingIndexes, nil
}

// markIndexesToPrint помечает индексы срок, окружающих выбранные
func markIndexesToPrint(indexes map[int]bool, after, before, context int, maxIndex int) {
	for i := range indexes {
		if context > 0 || before > 0 {
			markLinesBefore(i, max(context, before), indexes)
		}

		if context > 0 || after > 0 {
			markLinesAfter(i, max(context, after), indexes, maxIndex)
		}
	}
}

// markLinesBefore помечает before индексов срок перед данным индексом
func markLinesBefore(currentIndex, before int, indexes map[int]bool) {
	for i := currentIndex - before; i < currentIndex; i++ {
		if i > 0 {
			indexes[i] = true
		}
	}
}

// markLinesAfter помечает after индексов срок после данного индекса
func markLinesAfter(currentIndex, after int, indexes map[int]bool, maxIndex int) {
	for i := currentIndex; i < currentIndex+after+1; i++ {
		if i < maxIndex {
			indexes[i] = true
		}
	}
}

// max возвращает большее из чисел a и b.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
