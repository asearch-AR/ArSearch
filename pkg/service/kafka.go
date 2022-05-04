package service

import (
	"context"
	"strings"

	"github.com/segmentio/kafka-go"
)

func GetKafkaCli() (*kafka.Conn, error) {
	topic := "arsearch-topic"
	partition := 0

	conn, err := kafka.DialLeader(context.TODO(), "tcp", "localhost:9092", topic, partition)
	//conn, err := kafka.DialLeader(context.Background(), "tcp", "45.76.151.181:9092", topic, partition)
	return conn, err
}

func GetKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}