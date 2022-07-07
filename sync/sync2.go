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
func GetMirrorTxId1(offset int64) service_schema.ArData2 {

	res := service_schema.ArData2{}

	//api.viewblock.io/arweave/address/Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c?page=1
	url := "https://api.viewblock.io/arweave/addresses/Ky1c1Kkt-jZ9sY1hvLF5nCf6WWdBhIU5Un_BMYh-t3c?network=mainnet&page=%d"
	url = fmt.Sprintf(url,offset)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("err===>", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("If-None-Match", "W/\"9841-PZ35DpkNzBKubDpmzTcfZoiT8Vs\"")
	req.Header.Add("Origin", "https://viewblock.io")
	req.Header.Add("Referer", "https://viewblock.io/")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Site", "same-site")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	req.Header.Add("sec-ch-ua", "\".Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"103\", \"Chromium\";v=\"103\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")

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

	//fmt.Println(string(body))
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
		for _, v := range txIdList.Txs.Docs {
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
