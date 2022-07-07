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
	Contributor     string    `json:"contributor"`
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

	ArWeaveTx string `json:"arweave_tx"` //原始arweave_txid
}

type ArData2 struct {
	Hash string `json:"hash"`
	ChecksumAddress string `json:"checksumAddress"`
	Extra struct {
	} `json:"extra"`
	Balance string `json:"balance"`
	TxCount int `json:"txCount"`
	InternalsCount int `json:"internalsCount"`
	Txs struct {
		Docs []struct {
			Hash string `json:"hash"`
			BlockHeight int `json:"blockHeight"`
			Timestamp int64 `json:"timestamp"`
			Status string `json:"status"`
			Fee string `json:"fee"`
			Labels struct {
			} `json:"labels"`
			AddressTypes struct {
			} `json:"addressTypes"`
			Direction string `json:"direction"`
			From string `json:"from"`
			To string `json:"to"`
			Value string `json:"value"`
			Type string `json:"type"`
			Extra struct {
				Tags []struct {
					Name string `json:"name"`
					Value string `json:"value"`
				} `json:"tags"`
				Owner string `json:"owner"`
				LastTx string `json:"lastTx"`
				Dapp struct {
					Name string `json:"name"`
					Key string `json:"key"`
				} `json:"dapp"`
				DataSize string `json:"dataSize"`
				ObservedCount int `json:"observedCount"`
			} `json:"extra"`
		} `json:"docs"`
		Page string `json:"page"`
		Limit int `json:"limit"`
		Pages int `json:"pages"`
		Total int `json:"total"`
		Type string `json:"type"`
	} `json:"txs"`
}

type MirrorSearchRes struct {
	MirrorData
	ArweaveLink string `json:"arweave_link"`
	Score       float64
}
