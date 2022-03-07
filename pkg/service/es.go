package service

import (
	"ArSearch/pkg/service/service_schema"
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
)

var esClient *elastic.Client

func init() {
	esClient, _ = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200/"))
}

func PutToEs(article *service_schema.ArArticle) (string, error) {

	res, err := esClient.Index().
		Index("ar_search").
		Type("ar_article").
		Id(article.ID).
		BodyJson(article).
		Do(context.Background())

	if err != nil {
		return "err", err
	}

	return res.Result, nil
}

func SearchInEs(termQuery string) ([]*service_schema.ArArticle, error) {
	fmt.Println("start====>")

	searchResult, err := esClient.Search().
		Index("ar_search"). // search in index "twitter"
		Type("ar_article").
		//Query(termQuery).        // specify the query
		Pretty(true).          // pretty print request and response JSON
		Do(context.Background())     // execute

	if err != nil {
		return nil,err
	}
	fmt.Println("err===>",err.Error())

	r1, _ := json.Marshal(searchResult)
	fmt.Println("res===>",string(r1))

	return nil,nil
	//arList := make([]*service_schema.ArArticle, 0)
	//var article service_schema.ArArticle
	//for _, item := range searchResult.Each(reflect.TypeOf(article)) {
	//	arList = append(arList,item.(*service_schema.ArArticle))
	//}
	//
	//return arList,nil
}
