package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	brokerAddress := "localhost:9092"
	topic := "quickstart-events"
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress}, GroupID: "my-group", Topic: topic, MinBytes: 10e3, MaxBytes: 10e3,
	})

	defer reader.Close()

	fmt.Println("MinBytes: ", 10e3, ", topic: ", topic)
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Ошибка при чтении сообщения: %v", err)
		}
		fmt.Printf("Get message: key=%s, value=%s\n", string(msg.Key), string(msg.Value))
	}
}
