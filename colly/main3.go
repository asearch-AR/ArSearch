package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

func main22() {
	fileName := "txid.yaml"
	var result string

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	rows := ReadInput()
	q := AddUrl(rows)

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		result = result + e.Text + "\n"
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	q.Run(c)

	f.WriteString(result)

	log.Printf("Scraping done, Please check file %q for results\n", fileName)
}

func ReadInput() []string {
	// Read from file
	b, err := ioutil.ReadFile("data.csv") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	str := string(b) // convert content to a 'string'

	// split each row
	rows := strings.Split(str, "\n")
	return rows
}

func AddUrl(rows []string) *queue.Queue {
	Q, _ := queue.New(
		2, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000},
	)
	for _, url := range rows {
		//"https://arweave.net/koehrOAeK5Lpc860JoLo-Gc6ODiCv0JCSwqtR8UWYGY"
		pre := "https://arweave.net/"
		pre += url
		Q.AddURL(pre)
	}
	return Q
}
