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


//mirror data struct
type MirrorData1 struct {
	Data struct {
		Transactions struct {
			Edges []struct {
				Node struct {
					ID string `json:"id"`
				} `json:"node"`
			} `json:"edges"`
		} `json:"transactions"`
	} `json:"data"`
}

type ArData struct {
	Content struct {
		Body      string `json:"body"`
		Timestamp int    `json:"timestamp"`
		Title     string `json:"title"`
	} `json:"content"`
	Digest     string `json:"digest"`
	Authorship struct {
		Contributor         string `json:"contributor"`
		SigningKey          string `json:"signingKey"`
		Signature           string `json:"signature"`
		SigningKeySignature string `json:"signingKeySignature"`
		SigningKeyMessage   string `json:"signingKeyMessage"`
		Algorithm           struct {
			Name string `json:"name"`
			Hash string `json:"hash"`
		} `json:"algorithm"`
	} `json:"authorship"`
	Nft struct {
	} `json:"nft"`
	Version        string `json:"version"`
	OriginalDigest string `json:"originalDigest"`
}

type MirrorSearchRes struct {
	MirrorData
	Score float64
}