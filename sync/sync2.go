package main

import (
	"ArSearch/pkg/service"
	"ArSearch/pkg/service/service_schema"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

//sync mirror data
func GetMirrorTxId1(offset int64) service_schema.ArData1 {

	res := service_schema.ArData1{}

	url := fmt.Sprintf("https://api.viewblock.io/arweave/addresses/Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c/txs?page=%d&network=mainnet", offset)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("err===>", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://viewblock.io")
	req.Header.Set("Referer", "https://viewblock.io/")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36")
	req.Header.Set("Sec-Ch-Ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"101\", \"Google Chrome\";v=\"101\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("err===>", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("err====>", err)
		//return []byte{}
	}

	//fmt.Println("===>", string(body))
	json.Unmarshal(body, &res)
	return res
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
	//写入一下arweave_txid
	data.ArWeaveTx = txId
	return data
}

func main() {

	cli, _ := service.GetKafkaCli()

	var i int64 = 1
	wg := sync.WaitGroup{}

	for {
		txIdList := GetMirrorTxId1(i)
		for _, v := range txIdList.Docs {
			wg.Add(1)

			go func(txId string) {
				info := GetTxInfo(txId)
				marshal, _ := json.Marshal(info)
				cli.Write(marshal)

				fmt.Println("info===>",string(marshal))
				wg.Done()
			}(v.Hash)
		}

		fmt.Println("==========>",i)
		i++
	}

	wg.Wait()
}
