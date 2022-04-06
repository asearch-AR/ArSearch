package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	topic := "quickstart-events"
	partition := 0

	//conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	conn, err := kafka.DialLeader(context.Background(), "tcp", "45.76.151.181:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	//conn.SetWriteDeadline(time.Now().Add(10*time.Second))
	//_, err = conn.WriteMessages(
	//	kafka.Message{Value: []byte("one!")},
	//	kafka.Message{Value: []byte("two!")},
	//	kafka.Message{Value: []byte("three!")},
	//)

	messages, err1 := conn.WriteMessages(kafka.Message{
		Value: []byte("hello"),
	})

	if err1 !=nil {
		fmt.Println(err1)
	}

	fmt.Println(messages)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
