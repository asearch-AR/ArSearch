package main

import (
	"fmt"

	"github.com/everFinance/arsyncer"
	"github.com/everFinance/goar/types"
)

// sync all arweave tx
func main() {
	nullFilterParams := arsyncer.FilterParams{
		Tags: []types.Tag{
			{
				Name: "Content-Type",
				Value: "text/html",
			},
		},
	} // non-file params
	startHeight := int64(879220)
	arNode := "https://arweave.net"
	concurrencyNumber := 10 // runtime concurrency number, default 10
	s := arsyncer.New(startHeight, nullFilterParams, arNode, concurrencyNumber, 15)

	// run
	s.Run()

	// subscribe tx
	for {
		select {
		case sTx := <-s.SubscribeTxCh():
			// process synced txs
			fmt.Println("TX_ID====>",sTx[0].ID)
		}
	}
}