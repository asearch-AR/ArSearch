package service

import (
	"fmt"
	"testing"
)

func TestSearchInEs(t *testing.T) {
	es, err := SearchInEs("")

	if err !=nil{
		fmt.Println(err)
	}
	fmt.Println("1111")
	t.Log("===>",es)
}
