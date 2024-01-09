# Golang Web Scraper

This repository contains a simple web scraper written in Go. The scraper fetches and parses web pages, rendering JavaScript when necessary, and extracts the text content of the pages.

## Overview

The web scraper is started by the main function in the `main.go` file.

The `javascript.go` file contains a function named `RenderJavaScript`. This function takes a URL as an argument and uses Selenium to render JavaScript from the given URL. The function returns the HTML source of the page or an error.

The `scraper.go` file contains a function named `ScrapeWebPage`. This function takes a URL as an argument, sends a GET request to the URL, parses the response body as HTML, and returns the text content of the page or an error.

The `concurrency.go` file contains a function named `ConcurrentScrape`. This function takes a slice of URLs as an argument, concurrently scrapes each URL by calling the `ScrapeWebPage` function, and returns a map where the keys are the URLs and the values are the scraped data or an error message.
