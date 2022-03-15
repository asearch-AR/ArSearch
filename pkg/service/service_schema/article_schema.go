package service_schema

type ArArticle struct {
	ID             string  `json:"id"`
	Owner          string  `json:"owner"`
	Title          string  `json:"title"`
	ArticleContext string  `json:"article_context"`
	DataSize       float64 `json:"data_size"`
}

type ArSearchRes struct {
	Score       float64
	RedirectUrl string
	Article     ArArticle
}
