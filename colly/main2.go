package main

import (
	"ArSearch/pkg/service"
	"ArSearch/pkg/service/service_schema"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type TxList struct {
	TxId []string `yaml:"tx_id"`
}

func getUrlList(fileName string) TxList {
	data, err := ioutil.ReadFile(fileName)
	list := TxList{}

	if err != nil {
		fmt.Println("err1===>", err)
		return list
	}

	err1 := yaml.Unmarshal(data, &list)
	if err1 != nil {
		fmt.Println("err2", err1)
		return list
	}

	return list
}

func formatString(str string) string {
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.TrimSpace(str)
	return str
}

func main() {
	urlList := getUrlList("colly/txid.yaml")
	//urlList := getUrlList("colly/test.yaml")

	m1 := make(map[string]service_schema.ArArticle, 300)

	// Instantiate default collector
	c := colly.NewCollector(colly.AllowURLRevisit(),
		colly.MaxDepth(50),
	)

	// create a request queue with 2 consumer threads
	//threads 数过大会直接被kill 掉...
	q, _ := queue.New(
		5, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 100000}, // Use default queue storage
	)

	c.OnRequest(func(r *colly.Request) {
		//fmt.Println("visiting", r.URL)
	})

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		key := e.Request.URL.String()
		article := m1[key]
		key1 := strings.ReplaceAll(key, "https://", "")
		split := strings.Split(key1, "/")
		e.Text = formatString(e.Text)
		article.Title = e.Text
		article.ID = split[1]
		m1[key] = article

		//fmt.Println("title====>",e.Request.URL,e.Text)
	})

	c.OnHTML("p", func(e *colly.HTMLElement) {
		key := e.Request.URL.String()
		article := m1[key]
		e.Text = formatString(e.Text)
		article.ArticleContext = article.ArticleContext + e.Text
		m1[key] = article
		//response := e.Text
	})

	for _, v := range urlList.TxId {
		url := fmt.Sprintf("https://arweave.net/%s", v)
		q.AddURL(url)
	}

	// Consume URLs
	q.Run(c)

	marshal, _ := json.Marshal(m1)
	fmt.Println("===>res", string(marshal))
	//fmt.Println("===>res",m1)

	////put to es
	for _, v := range m1 {
		if v.Title == "" || v.ID == "" || v.ArticleContext == "" {
			continue
		}
		es, err := service.PutToEs(&v)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(es)
	}

}