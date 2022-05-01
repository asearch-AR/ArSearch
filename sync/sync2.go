package main

import (
	"ArSearch/pkg/service"
	"ArSearch/pkg/service/service_schema"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/labstack/gommon/log"
)

//sync mirror data

func GetMirrorTxId() service_schema.MirrorData1 {
	v := service_schema.MirrorData1{}

	url := "https://arweave.net/graphql"

	payload := strings.NewReader("{\"query\":\"query {\\n    transactions(\\n        tags: [\\n            {\\n                name: \\\"App-Name\\\",\\n                values: [\\\"MirrorXYZ\\\"]\\n            }\\n        ]\\n    ) {\\n        edges {\\n            node {\\n                id\\n            }\\n        }\\n    }\\n}\",\"variables\":{}}")

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPost, url, payload)

	if err != nil {
		fmt.Println(err)
		return v
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return v
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return v
	}
	err1 := json.Unmarshal(body, &v)

	if err1 != nil {
		fmt.Println(err1)
		return v
	}

	return v
}

func GetTxInfo(txId string) service_schema.ArData {
	data := service_schema.ArData{}

	url := fmt.Sprintf("https://arweave.net/%s", txId)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		fmt.Println(err)
		return data
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return data
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return data
	}

	json.Unmarshal(body, &data)
	return data
}

func main() {
	cli, _ := service.GetKafkaCli()
x:
	ids := GetMirrorTxId()

	edges := ids.Data.Transactions.Edges

	wg := sync.WaitGroup{}
	for _, v := range edges {
		wg.Add(1)
		txId := v.Node.ID
		//2.fetch tx info
		go func() {
			info := GetTxInfo(txId)
			marshal, _ := json.Marshal(info)
			write, _ := cli.Write(marshal)
			log.Info("write===>", write)
			wg.Done()
		}()
	}
	wg.Wait()
	goto x
}
