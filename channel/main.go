package main

import (
	"errors"
	"fmt"
	"time"
)

func producer(ch chan<- int,err error) {
	defer func() {
		err1:=recover()
		if err!=nil {
			fmt.Println(err1)
			ch<--1
		}
	}()
	if err!=nil{
		panic(err)
	}
	time.Sleep(time.Second*5)
	ch<-100
}

func main() {
	ch := make(chan int, 100)

	go producer(ch,nil)
	go producer(ch,nil)
	go producer(ch,errors.New("err"))

	t1 := time.Now()
	//fmt.Println(len(ch))
	for i:=0;i<3;i++{
		v,ok:=<-ch
		if ok{
			fmt.Println(v)
		}else {
			fmt.Println("nil")
		}
	}

	t2 := time.Now()
	fmt.Println(t2.Sub(t1))
}
