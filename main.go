func main() {
	// This is the entry point of the application.
	// The main function initializes and starts the web scraper.
	startWebScraper()
}

import (
	"net/http"
	"html"
	"sync"
	"github.com/tebeka/selenium"
)
