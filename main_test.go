package main

import (
	"fmt"
	"sync"

	"io/ioutil"
	"net/http"
	"testing"
)

func DoGet(){
	url := "http://localhost:9999/query_mirror?q=mirror"
	//url := "https://api.asearch.io/search_mirror?q=hello"
	method := "GET"

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body[1]))
}

func Benchmark_main(b *testing.B) {
	wg:=sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			DoGet()
			wg.Done()
		}()
	}
	wg.Wait()
}
