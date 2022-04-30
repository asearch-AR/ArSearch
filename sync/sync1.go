package main

import (
	"encoding/json"
	"fmt"
	"github.com/everFinance/arsyncer"
	"github.com/everFinance/goar/types"
)

// sync all arweave tx
func main() {
	nullFilterParams := arsyncer.FilterParams{
		Tags: []types.Tag{
			{
				Name: "App-Name",
				Value: "MirrorXYZ",
			},
		},
	} // non-file params
	startHeight := int64(0)
	arNode := "https://arweave.net"
	concurrencyNumber := 10 // runtime concurrency number, default 10
	s := arsyncer.New(startHeight, nullFilterParams, arNode, concurrencyNumber, 15)

	// run
	s.Run()

	// subscribe tx
	for {
		select {
		case sTx := <-s.SubscribeTxCh():

			for _,v:=range sTx{
				marshal, _ := json.Marshal(v)

				fmt.Println("===>",string(marshal))
			}

		}
	}
}