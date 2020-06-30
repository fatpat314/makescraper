package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	// "strings"

	"github.com/gocolly/colly"
)

type Count struct{
	TotalCases string `json:"Total Cases"`
	TotalDeaths string `json:"Total Deaths"`

}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/


func scraper(w http.ResponseWriter, r *http.Request) {
	// Instantiate default collector
	c := colly.NewCollector()

	selector := "body > div.container.d-flex.flex-wrap.body-wrapper.bg-white > main > div:nth-child(3) > div > div.syndicate > div:nth-child(1) > div > div > div > section > div > div > div:nth-child(1)"
	// d := strings.Replace(structure, "body > div.container.d-flex.flex-wrap.body-wrapper.bg-white > main > div:nth-child(3) > div > div.syndicate > div:nth-child(1) > div > div > div > section > div > div > div:nth-child(1) > span.count", "body > div.container.d-flex.flex-wrap.body-wrapper.bg-white > main > div:nth-child(3) > div > div.syndicate > div:nth-child(1) > div > div > div > section > div > div > div:nth-child(2) > span.count", -1)

	// On every a element which has href attribute call callback																																	 vChange this to change the amount
	c.OnHTML(selector, func(e *colly.HTMLElement) {
    	link := e.Attr("span.count")
		// totalCases := strings.Replace(e.Text, )
		fmt.Println()

		count := Count{
			TotalCases: e.Text,
			// TotalDeaths: e.Text,
		}



		// Print link
    	fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		returnJSON, err := json.Marshal(count)
		if err != nil {
			fmt.Println("error")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(returnJSON)

		fmt.Println(returnJSON)
	})


	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.cdc.gov/coronavirus/2019-ncov/cases-updates/cases-in-us.html")
}

func main(){
	host := "0.0.0.0:8888"
	http.HandleFunc("/", scraper)
	fmt.Println("Starting server: http//" + host)
	err := http. ListenAndServe(host, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// scraper()
}
