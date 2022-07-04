package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func main() {

	c := colly.NewCollector(
		colly.UserAgent("User-Agent: Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0"),
		colly.AllowURLRevisit(),
	)

	var proxyList []Proxy
	c.OnHTML(".list-proxy", func(e *colly.HTMLElement){

		// check lastUpdated first and prevent insert 
		tbody := e.DOM.ChildrenMatcher(goquery.Single("tbody"))
		rows := tbody.ChildrenMatcher(goquery.Single("tr"))

		rows.Each(func(rowIndex int, row *goquery.Selection) {
			cells := row.ChildrenMatcher(goquery.Single("td"))

			values := cells.Map(func(cellIndex int, cell *goquery.Selection) string {
				return cell.Text()
			})

			newProxy := Proxy{
				Address: strings.Trim(values[0], ""),
				Port: strings.Trim(values[1], ""),
				Country: strings.Trim(values[2], ""),
				RegionOrCity: strings.Trim(values[3], ""),
				Protocol: strings.Trim(values[4], ""),
				Anon: strings.Trim(values[5], ""),
				Speed: strings.Trim(values[6], ""),
			}
			proxyList = append(proxyList, newProxy)


			log.Printf("%+v", newProxy)
		})
	})

	c.OnResponse(func(res *colly.Response) {
		log.Printf("Response: %+v", res.StatusCode)
	})

	c.OnRequest(func(req *colly.Request) {
		log.Printf("Request to %s", req.URL.String())
	})

	c.Visit("https://freeproxyupdate.com/socks5-proxy")


	// cleanup
	// test each proxy to make sure is working at the time of the insert (could insert anyway - maybe comes online before usage)
	// persistence - google sheet or db
	// put on romarg vps


	// stretch -  scrap more free lists, give spys another try
}
