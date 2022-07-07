package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"ArSearch/pkg/service"
	"ArSearch/pkg/service/service_schema"
)

func fetch(url string) []byte {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Println(err)
		return []byte{}
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return []byte{}
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return []byte{}
	}
	return body
}

const url = "https://mirror.cjpais.com/api/publisher/%s"

func main11() {
	//1.fetch from https://mirror.cjpais.com/api/publishers
	url1 := "https://mirror.cjpais.com/api/publishers"
	publisherBytes := fetch(url1)
	publishers := make([]service_schema.Publisher, 0)
	json.Unmarshal(publisherBytes, &publishers)

	//2.fetch mirror raw data
	for _, publisher := range publishers {
		url2 := fmt.Sprintf(url, publisher.Name)
		fmt.Println(url2)
		bytes := fetch(url2)
		mirrorData := make([]service_schema.MirrorData, 0)
		json.Unmarshal(bytes, &mirrorData)

		wg := sync.WaitGroup{}
		for _, v := range mirrorData {
			wg.Add(1)
			go func() {
				data, err := service.SaveMirrorData(&v)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(data)
				wg.Done()
				time.Sleep(time.Millisecond * 200)
			}()
		}

		wg.Wait()
	}
}
