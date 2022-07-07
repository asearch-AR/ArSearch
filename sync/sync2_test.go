package main

import (
	"fmt"
	"testing"
)

func TestGetMirrorTxId1(t *testing.T) {
	id1 := GetMirrorTxId1(1)
	fmt.Println(id1)
}

func Test_GetTxInfo(t *testing.T){
	info := GetTxInfo("GN6u08fVddJD_uB0IPWCh7AcfvFdZ35c9bE7EM8bFC8")
	fmt.Println(info)
}