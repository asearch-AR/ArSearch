package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type Ctx1 struct {
	Title string
	Ctx 	[]string
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		//colly.AllowedDomains("arweave.net"),
		colly.MaxDepth(20),
	)

	m := make(map[string]Ctx1, 1)

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		m[e.Request.URL.String()] = Ctx1{Title: e.Text}
		//fmt.Println("====>",e.Text)
		//fmt.Println(e.Request.URL)
		//fmt.Println(e.Text)
	})

	// On every a element which has href attribute call callback
	c.OnHTML("p", func(e *colly.HTMLElement) {
		ctx1 := m[e.Request.URL.String()]
		ctx1.Ctx = append(ctx1.Ctx,e.Text)
		//fmt.Println("----->",ctx1.Ctx)
		fmt.Println("=======>",ctx1)
		m[e.Request.URL.String()] = ctx1
		//link := e.Attr("href")
		// Print link
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		//m[e.Request.URL.String()] = Ctx1{Title: e.Text}
		//fmt.Println("====>",e.Text)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		//c.Visit(e.Request.AbsoluteURL(link))
	})



	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})


	//拼url :https://arweave.net/[tx_id]
	// Start scraping on https://hackerspaces.org
	//c.Visit("https://arweave.net/-Vh8fzPNQTJ7UP_h2cnfSY47Zfyo5LhqTqrm5IOivI8")
	c.Visit("https://mirror.xyz/0x707D306714FF28560f32bF9DAE973BD33cd851c5/96oLZtaMi1XgialvnpEsk7kEwen3Daqy7PBWOb0jItk")
	//c.Visit("https://arweave.net/koehrOAeK5Lpc860JoLo-Gc6ODiCv0JCSwqtR8UWYGY")

	fmt.Println("res===>",m)
}