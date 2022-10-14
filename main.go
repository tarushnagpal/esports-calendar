package main

import (
	"fmt"
	"time"

	"github.com/CallMeTarush/esports-calendar/webscraper/dota"
)

func main() {
	currentTime := time.Now()
	for i := 0; i <= 30; i++ {
		date := currentTime.AddDate(0, 0, i).Format("2006-01-02")
		fmt.Println("Scraping" + date)
		webscraper.ScrapeGameflow(date)
	}
}