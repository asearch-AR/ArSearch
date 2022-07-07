package main

import (
	"fmt"
	"testing"
)

func TestGetMirrorTxId1(t *testing.T) {
	id1 := GetMirrorTxId1(1)
	fmt.Println(id1)
}
