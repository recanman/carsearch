// made by recanman
package scraper

import (
	"carsearch/pkg/models"
	"fmt"
	"os"
	"time"
)

var Notifications = make(chan models.Listing)
var Failure = make(chan error)

func addScrape(id string, body []byte) error {
	// makedirectory scrapes
	os.Mkdir("scrapes", 0755)

	//current timestamp + id
	fileName := fmt.Sprintf("scrapes/scrape-%d-%s.html", time.Now().Unix(), id)
	return os.WriteFile(fileName, body, 0644)
}
