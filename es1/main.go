package main

import (
	"context"
	"fmt"

	"github.com/olivere/elastic"
)

type Tmp struct {
	Name string
	ID   int64
}

//存储编辑文章
func main(){

	esClient, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200/"))

	if err != nil {
		fmt.Println("connect es error", err.Error())
	}

	fmt.Println("start====>")

	_,err1 := esClient.Search().
		Index("ar_search"). // search in index "twitter"
		Type("ar_article").
		//Query(termQuery).        // specify the query
		//Pretty(true).          // pretty print request and response JSON
		Do(context.Background())     // execute

	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println("err===>",err1.Error())




	//t := new(Tmp)
	//t.Name = "guruiqin"
	//t.ID   = 1234
	//
	//res, err1 := client.Index().
	//	Index("test").
	//	Type("test").
	//	Id("1234").
	//	BodyJson(t).
	//	Do(context.Background())
	//
	//if err1 != nil {
	//	fmt.Println("err==>",err1)
	//}
	//
	//fmt.Println("success",res.Result)
	//fmt.Println("res",putRes.Status)
}

////获取博客列表
//func BlogList(pageNum int) []model.Article {
//
//	if es == nil {
//		log.Println("BlogList es nil")
//		return nil
//	}
//
//	searchResult, err := es.Search().
//		Index(constant.EsIndexBlog). // search in index "twitter"
//		// Query(termQuery).        // specify the query
//		Sort("article_id", true).    // sort by "user" field, ascending
//		From(pageNum * 10).Size(10). // take documents 0-9
//		Pretty(true).                // pretty print request and response JSON
//		Do(context.Background())     // execute
//	if err != nil {
//		// Handle error
//		panic(err)
//	}
//
//	blogList := make([]model.Article, 10)
//	var article model.Article
//	for index, item := range searchResult.Each(reflect.TypeOf(article)) {
//		blogList[index] = item.(model.Article)
//	}
//
//	return blogList
//}
//
//// 通过文章标题与标签搜索文章
//func SearchBlogByTitleAndTags() {
//
//	if es == nil {
//		return
//	}
//
//}