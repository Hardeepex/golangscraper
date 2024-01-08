package main

import (
	"net/http"
	"golang.org/x/net/html"
	"strings"
	"sync"
)

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

func ConcurrentScrape(urls []string) map[string]string {
	var wg sync.WaitGroup
	results := make(map[string]string)

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			data, err := ScrapeWebPage(url)
			if err != nil {
				results[url] = err.Error()
			} else {
				results[url] = data
			}
		}(url)
	}

	wg.Wait()
	return results
}
