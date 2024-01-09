package main

import (
	"net/http"
	"golang.org/x/net/html"
	"strings"
	"sync"
)

// ScrapeWebPage scrapes the given URL's web page and returns the text content.
func ScrapeWebPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	var f func(*html.Node)
	var sb strings.Builder
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			sb.WriteString(n.Data)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return sb.String(), nil
}

func main() {
	var wg sync.WaitGroup
	urls := []string{"http://example.com", "http://example.org", "http://example.net"}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			data, err := ScrapeWebPage(url)
			if err != nil {
				panic(err)
			}
			println(data)
		}(url)
	}
	wg.Wait()
}
