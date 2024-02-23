package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

// runWget запускает утилиту.
func runWget(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		writeError(cmd.Usage())
	}

	recursive, err := cmd.Flags().GetBool("recursive")
	if err != nil {
		writeError(err)
	}

	url := cleanUrl(args[0])
	if recursive {
		downloadPages(url)
	} else {
		downloadPage(url)
	}
}

// downloadPage скачивает index.html по адресу url.
func downloadPage(url string) error {
	body, err := getPageBody(url)
	if err != nil {
		return err
	}
	defer body.Close()

	safePage("index.html", body)

	return nil
}

// downloadPages скачивает все страницы сайта  url.
func downloadPages(url string) {
	var download func(url, path string)
	visited := make(map[string]bool, 0)

	download = func(url, path string) {
		cleanQuery(&path)
		cleanHash(&path)
		addForwardSlash(&path)

		if visited[path] {
			return
		}
		visited[path] = true

		body, err := getPageBody(url + path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		var buf bytes.Buffer
		tee := io.TeeReader(body, &buf) // чтобы прочитать из body дважды

		err = os.MkdirAll(url+path, os.ModePerm)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		err = safePage(url+path+"/index.html", tee)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		fmt.Fprintf(os.Stdout, "downloaded: %s%s\n", url, path)

		for _, link := range getLinks(&buf) {
			// подходят только относительные пути
			if !hasDot(link) {
				download(url, link)
			}
		}
	}

	download(url, "/")
}
