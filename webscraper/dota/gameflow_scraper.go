package webscraper

import (
	"log"
	"strings"

	"github.com/CallMeTarush/esports-calendar/util"
	"github.com/CallMeTarush/esports-calendar/googlecalendar"

	"github.com/gocolly/colly/v2"
)

func ScrapeGameflow(matchDate string) {
	log.Println("Scraping Game Flow")

	addedMatches := make([]string, 200);

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		queryParams := util.GetQueryParams(link)
		queryParamLength := len(queryParams)
		
		isMatchLink := false
		var seriesId string

		if queryParamLength > 0 {
			lastQueryParam := queryParams[queryParamLength - 1]
			if strings.Contains(lastQueryParam, "seriesId") {
				seriesId = strings.SplitAfter(lastQueryParam, "=")[1]
				isMatchLink = true
			}
		}

		if !isMatchLink {
			return
		}
		
		teamNames := make([]string, 2)
		var time string

		// log.Println("Check", e.childText)
		teamNameFilled := 0
		// Check each div for number of p tags and img tags
		e.ForEach("div", func(_ int, el *colly.HTMLElement) {

			// If there is not exactly 1 p tag, exit.
			if el.DOM.Find("p").Length() != 1 {
				return;
			}

			// If an img tag was found then its a team name
			if el.DOM.Find("img").Length() == 1 {
				log.Println(el.ChildText("p"))
				teamNames[teamNameFilled] = el.ChildText("p")
				teamNameFilled = 1
			} else {
				// If no img tag then its the time
				time = el.ChildText("p")
			}

		})
		log.Println(time)
		if len(seriesId) > 0 && !util.ArrayContains(addedMatches, seriesId) {
			addedMatches = append(addedMatches, seriesId)
			googlecalendar.AddGoogleCalendarEvent(teamNames, time, matchDate)
		}
		log.Println(addedMatches)

	})

	c.Visit("https://gameflow.tv/dota/matches?date=" + matchDate + "T00%3A00%3A00%2B05%3A30")
}