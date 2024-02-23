package cmd

import (
	"fmt"
	"io"
	"net/http"
)

// getPageBody возвращает body страницы по url.
func getPageBody(url string) (io.ReadCloser, error) {
	resp, err := http.Get("http://" + url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d", http.StatusOK)
	}

	return resp.Body, nil
}
