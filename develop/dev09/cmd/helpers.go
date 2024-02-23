package cmd

import (
	"fmt"
	"os"
	"strings"
)

// addForwardSlash добавляет / в начало пути.
func addForwardSlash(path *string) {
	if !strings.HasPrefix(*path, "/") {
		*path = "/" + *path
	}
}

// cleanHash отчищает путь от параметров запроса.
func cleanQuery(path *string) {
	index := strings.Index(*path, "?")
	if index != -1 {
		*path = (*path)[:index]
	}
}

// cleanHash отчищает путь от якоря.
func cleanHash(path *string) {
	index := strings.Index(*path, "#")
	if index != -1 {
		*path = (*path)[:index]
	}
}

// hasDot проверяется, есть ли в пути точка
func hasDot(link string) bool {
	return strings.Contains(link, ".")
}

// cleanUrl убирает из url протокол и trailing slash
func cleanUrl(url string) string {
	return strings.TrimSuffix(strings.Split(url, "://")[1], "/")
}

// writeError выводит err в stderr.
func writeError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
