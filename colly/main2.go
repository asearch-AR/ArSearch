package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"ArSearch/pkg/service"
	"ArSearch/pkg/service/service_schema"

	"gopkg.in/yaml.v2"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

type TxList struct {
	TxId []string `yaml:"tx_id"`
}

func getUrlList(fileName string) TxList{
	data, err := ioutil.ReadFile(fileName)
	list := TxList{}

	if err!=nil{
		fmt.Println("err1===>",err)
		return list
	}

	err1 := yaml.Unmarshal(data, &list)
	if err1!=nil{
		fmt.Println("err2",err1)
		return list
	}

	return list
}

func formatString(str string)string{
	str =strings.ReplaceAll(str,"\n","")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.TrimSpace(str)
	return str
}

func main222() {
	urlList := getUrlList("colly/txid.yaml")

	m1 := make(map[string]service_schema.ArArticle, len(urlList.TxId))

	// Instantiate default collector
	c := colly.NewCollector(colly.AllowURLRevisit())

	// create a request queue with 2 consumer threads
	q, _ := queue.New(
		100, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		key:=e.Request.URL.String()
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


	for _,v:=range urlList.TxId{
		url := fmt.Sprintf("https://arweave.net/%s", v)
		q.AddURL(url)
	}

	// Consume URLs
	q.Run(c)


	//put to es
	for _,v:=range m1{
		if v.Title == "" || v.ID == "" || v.ArticleContext == ""{
			continue
		}

		es, _ := service.PutToEs(&v)
		fmt.Println(es)
	}

}