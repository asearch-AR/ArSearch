package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"

	"ArSearch/pkg/service/service_schema"
)

var es *elasticsearch.Client

func init() {

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://45.76.151.181:9200",
			"http://45.76.151.181:9201",
		},
	}

	es, _ = elasticsearch.NewClient(cfg)
}

func PutToEs(article *service_schema.ArArticle) (string, error) {
	m1, _ := json.Marshal(article)

	req := esapi.IndexRequest{
		Index:        "ar_search",
		DocumentType: "ar_article",
		DocumentID:   article.ID,
		Body:         strings.NewReader(string(m1)),
		Refresh:      "true",
	}

	res, err := req.Do(context.Background(), es)

	if err != nil {
		return "", err
	}

	return res.String(), nil
}

func SearchInEs(termQuery string) ([]service_schema.ArSearchRes, error) {

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool":map[string]interface{}{
				"should":[]map[string]interface{}{
					{"match": map[string]interface{}{"article_context":termQuery,}},
					{"match": map[string]interface{}{"title":termQuery,}},
				},
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	fmt.Println("query==>",buf.String())


	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("ar_search"),
		es.Search.WithDocumentType("ar_article"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)

	if err != nil {
		return []service_schema.ArSearchRes{}, err
	}
	defer res.Body.Close()

	var r map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	searchResList := make([]service_schema.ArSearchRes, 0)
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		article := hit.(map[string]interface{})["_source"].(map[string]interface{})
		score := hit.(map[string]interface{})["_score"]
		res := service_schema.ArSearchRes{
			Score: score.(float64),
			Article: service_schema.ArArticle{
				ID:             article["id"].(string),
				ArticleContext: article["article_context"].(string),
				Title:          article["title"].(string),
			},
			RedirectUrl: fmt.Sprintf("https://arweave.net/%s",article["id"].(string)),
		}


		searchResList = append(searchResList, res)
	}

	return searchResList, nil
}
