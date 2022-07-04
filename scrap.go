package main

import (
	"log"

	"github.com/gocolly/colly"
)

const proxyList = "https://spys.one/en/socks-proxy-list/"



func scrap() {

	c := colly.NewCollector(
		colly.AllowedDomains("spys.one"),
	)

	c.OnHTML("font.spy14", func(e *colly.HTMLElement){
		log.Println(e.Text)
	})

	c.OnResponse(func(res *colly.Response) {
		log.Printf("Response: %+v", res.StatusCode)
	})

	c.OnRequest(func(req *colly.Request) {
		log.Printf("Request to %s", req.URL.String())
	})

	c.Visit(proxyList)

}



/*
func scrap() {
	c := colly.NewCollector()

	requestData := url.Values{}
	requestData.Set("xx0", "e2852f028f85c230e62a64024eb3090b")
	requestData.Set("xpp", "5")
	requestData.Set("xf1", "0")
	requestData.Set("xf2", "0")
	requestData.Set("xf4", "0")
	requestData.Set("xf5", "2")
	encodedData := requestData.Encode()
	requestBodyReader := strings.NewReader(encodedData)

	headerC := http.Header{}

	headerC.Add(http.CanonicalHeaderKey("Content-Type"), "application/x-www-form-urlencoded")
	headerC.Add(http.CanonicalHeaderKey("Cookie"), "__cf_bm=QR720UUoMktsOmcrYv3E2_xVZQnwL_lVERdoMAmnMTw-1656957669-0-AZ1/PZEJ7kDgybkNsdlJZxxS40YFrufUKVwea0CeuJpTeum55FjAPbbRb3hn1F3PQvX1bnMjEI8wyqY8/yPdT+odwieoWTK4TwBJEGqUEOt7Ig56RXyij2+u5NRDR4JXVA==; cf_chl_2=c6aabe5fa46af00; cf_chl_prog=x10; cf_clearance=Yjp3uu8QMfqN9V5o.f.ltrWoByyWJpDUcSiJx.Jq3Xw-1656957675-0-150")

	headerC.Add(http.CanonicalHeaderKey("host"), "spys.one")
	headerC.Add(http.CanonicalHeaderKey("origin"), "https://spys.one")
	headerC.Add(http.CanonicalHeaderKey("User-Agent"), "Mozilla/5.0 (X11; Linux x86_64; rv:102.0) Gecko/20100101 Firefox/102.0")

	headerC.Add(http.CanonicalHeaderKey("Accept-Language"), "en-US,en;q=0.5") 
	headerC.Add(http.CanonicalHeaderKey("Accept-Encoding"), "gzip, deflate, br") 
	headerC.Add(http.CanonicalHeaderKey("Referer"), "https://spys.one/en/socks-proxy-list/") 
	headerC.Add(http.CanonicalHeaderKey("Alt-Used"), "spys.one") 
	headerC.Add(http.CanonicalHeaderKey("Sec-Fetch-Dest"), "document") 
	headerC.Add(http.CanonicalHeaderKey("Sec-Fetch-Mode"), "navigate") 
	headerC.Add(http.CanonicalHeaderKey("Sec-Fetch-Site"), "same-origin") 
	headerC.Add(http.CanonicalHeaderKey("TE"), "trailers") 
	log.Println(headerC)


	c.OnHTML("a", func(e *colly.HTMLElement) {
		log.Println(e.Text)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println(r.StatusCode)

	})

	err := c.Request("POST", proxyList, requestBodyReader, nil, headerC)
	if err != nil {
		log.Fatalf("Failed to request %+v", err)
	}

	//c.Visit(proxyList)
}
*/
