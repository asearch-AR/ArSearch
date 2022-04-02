package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		//colly.AllowedDomains("arweave.net"),
		colly.MaxDepth(10),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("p", func(e *colly.HTMLElement) {
		//link := e.Attr("href")
		// Print link
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		fmt.Println("====>",e.Text)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		//c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	//拼url :https://arweave.net/[tx_id]
	// Start scraping on https://hackerspaces.org
	c.Visit("https://arweave.net/koehrOAeK5Lpc860JoLo-Gc6ODiCv0JCSwqtR8UWYGY")
}