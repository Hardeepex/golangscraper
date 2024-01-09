package main

import (
	"net/http"
	"github.com/tebeka/selenium"
	"golang.org/x/net/html"
)

// RenderJavaScript renders JavaScript from a given URL using Selenium and returns the resulting HTML source.
func RenderJavaScript(url string) (string, error) {
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		return "", err
	}
	defer wd.Quit()

	err = wd.Get(url)
	if err != nil {
		return "", err
	}

	html, err := wd.PageSource()
	if err != nil {
		return "", err
	}

	return html, nil
}
