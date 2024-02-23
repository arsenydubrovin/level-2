package cmd

import (
	"io"

	"golang.org/x/net/html"
)

// getLinks возвращает список всех ссылок из тегов <a> в body.
func getLinks(body io.Reader) []string {
	links := make([]string, 0)

	tokenizer := html.NewTokenizer(body)
	for {
		tokenType := tokenizer.Next()

		if isLastToken(tokenType) {
			return links
		}

		if !isTagToken(tokenType) {
			continue
		}

		token := tokenizer.Token()
		if token.Data != "a" {
			continue
		}

		if link, ok := fetchLink(token); ok {
			links = append(links, link)
		}
	}
}

// fetchLink извлекает значени аттрибута href из токена.
func fetchLink(t html.Token) (string, bool) {
	for _, attr := range t.Attr {
		if attr.Key == "href" {
			return attr.Val, true
		}
	}
	return "", false
}

// isTagToken проверят, является ли токен открывающим или закрывающим тегом.
func isTagToken(tt html.TokenType) bool {
	return tt == html.StartTagToken || tt == html.EndTagToken || tt == html.SelfClosingTagToken
}

// isLastToken проверяет, является ли токен последним на странице.
func isLastToken(tt html.TokenType) bool {
	return tt == html.ErrorToken
}
