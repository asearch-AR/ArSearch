package service_schema

type ArArticle struct {
	ID             string  `json:"id"`
	Title          string  `json:"title"`
	ArticleContext string  `json:"article_context"`
}

type ArSearchRes struct {
	Score       float64
	RedirectUrl string
	Article     ArArticle
}
