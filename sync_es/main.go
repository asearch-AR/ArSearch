package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main()  {
	ctx:=context.TODO()
	group, _ := errgroup.WithContext(ctx)

	for {
		group.Go(func() error {
			fmt.Println("====>")
			return nil
		})
	}

	group.Wait()
}

