package service_schema

type ArArticle struct {
	ID             string `json:"id"`
	Owner          string `json:"owner"`
	Title          string `json:"title"`
	ArticleContext string `json:"article_context"`
	DataSize       int64  `json:"data_size"`
}

//