package service_schema

import "time"

type ArArticle struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	ArticleContext string `json:"article_context"`
}

type Publisher struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"createdAt"`
}

type MirrorData struct {
	Id              int64     `json:"id"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	CreatedAt       time.Time `json:"createdAt"`
	PublishedAt     time.Time `json:"publishedAt"`
	Digest          string    `json:"digest"`
	Link            string    `json:"link"`
	OriginalDigest  string    `json:"originalDigest"`
	PublicationName string    `json:"publicationName"`
	Cursor          string    `json:"cursor"`
	ArweaveTx       string    `json:"arweaveTx"`
	BlockHeight     int64     `json:"blockHeight"`
}

type ArSearchRes struct {
	Score       float64
	RedirectUrl string
	Article     ArArticle
}

type MirrorSearchRes struct {
	MirrorData
	Score float64
}