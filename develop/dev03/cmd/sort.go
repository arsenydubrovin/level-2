package cmd

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// runSort запускает сортировку строк с переданными агрументами
func runSort(cmd *cobra.Command, args []string) {
	rowLines, err := readLines(args[0])
	if err != nil {
		fmt.Println("error reading file:", err)
		return
	}

	if uniqueLines {
		rowLines = removeDuplicates(rowLines)
	}

	lines := splitLines(rowLines)

	sort.SliceStable(lines, func(i, j int) bool {
		return compareLines(lines[i], lines[j], numericSort, sortColumn)
	})

	if reverseOrder {
		slices.Reverse(lines)
	}

	printLines(lines)
}

// readLines читает строки файла по filePath
func readLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// removeDuplicates возвращает слайс без повторяющийхся строк
func removeDuplicates(lines []string) []string {
	uniqueLines := make([]string, 0)
	set := make(map[string]struct{})

	for _, line := range lines {
		if _, duplicate := set[line]; !duplicate {
			uniqueLines = append(uniqueLines, line)
		}

		set[line] = struct{}{}
	}

	return uniqueLines
}

// splitLines разделяет все строки переданного на слова по пробелу
func splitLines(lines []string) [][]string {
	splittedLines := make([][]string, 0, len(lines))

	for _, line := range lines {

		// пустая строка долна дозавляться как пустой слайс
		if line == "" {
			splittedLines = append(splittedLines, []string{""})
			continue
		}

		splittedLines = append(splittedLines, strings.Fields(line))
	}

	return splittedLines
}

// compareLines сравнивает два слайса строк по элементу с номером col (нумерация с 1)
func compareLines(line1, line2 []string, byNumber bool, col int) bool {
	if col <= 0 || col > len(line1) || col > len(line2) {
		col = 1 // если элемента с номером col нет, сравнение идёт по первому элементу
	}

	return compareStrings(line1[col-1], line2[col-1], byNumber)
}

// compareStrings сравнивает меджу собой две строки. Если первая строка меньше второй, возвращает true.
//
// Если byNumber = true, сравнение идёт не в алфавитном порядке, а с учётом номера в начале строки
func compareStrings(str1, str2 string, byNumber bool) bool {
	if byNumber {
		n1, hasNumber1 := extractDigitalPrefix(str1)
		n2, hasNumber2 := extractDigitalPrefix(str2)

		if !hasNumber1 && !hasNumber2 {
			return str1 < str2 // если у обеих строк нет чисел в начале, сравнивнение идёт в алфавитном порядке
		}

		return n1 < n2
	}

	return str1 < str2
}

// extractDigitalPrefix извлекает число из начала переданной строки.
//
// Возвращает true, если у строки есть числовой префикс. Если строка пустая, возвращает -Inf и false. Если числа в начале строки нет, возвращает +Inf и false.
func extractDigitalPrefix(str string) (float64, bool) {
	if str == "" {
		return math.Inf(-1), false // пустая строка стоит перед всеми числами
	}

	re := regexp.MustCompile(`^[-]?[0-9]*\.?[0-9]+`)

	match := re.FindString(str)

	if match != "" {
		number, _ := strconv.ParseFloat(match, 64)
		return number, true
	}

	return math.Inf(1), false // строка без числа в начале стоит после всех чисел
}

// printLines печатает отсортированные строки в stdout.
func printLines(lines [][]string) {
	for _, line := range lines {
		fmt.Fprintln(os.Stdout, strings.Join(line, " "))
	}
}
