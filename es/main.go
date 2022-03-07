package main

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
			"http://localhost:9201",
		},
		// ...
	}

	es, _ := elasticsearch.NewClient(cfg)
	log.Println(elasticsearch.Version)
	log.Println(es.Info())

	var b strings.Builder
	b.WriteString(`{"title" : "`)
	b.WriteString("title")
	b.WriteString(`"}`)

	fmt.Println(b.String())
	fmt.Println(strings.NewReader(b.String()))
	req := esapi.IndexRequest{
		//Index:      "http:localhost:9200/test",
		Index:      "test",
		DocumentID: "1234",
		Body:       strings.NewReader(b.String()),
		Refresh:    "true",
	}

	// Perform the request with the client.
	res, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	fmt.Println("res===>",res.String())
}