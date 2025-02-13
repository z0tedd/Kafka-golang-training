package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type MyMessage struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func main() {
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatalf("failed to create producer: %v", err)
	}
	defer producer.Close()
	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer consumer.Close()
	partConsumer, err := consumer.ConsumePartition("ping", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer partConsumer.Close()
	for msg := range partConsumer.Messages() {
		var receivedMessage MyMessage
		err := json.Unmarshal(msg.Value, &receivedMessage)
		if err != nil {
			log.Print(err)
			continue
		}
		log.Printf("Received message: %+v\n", receivedMessage)
		responseText := fmt.Sprintf("%s %s (%s)", receivedMessage.Name, receivedMessage.Value, receivedMessage.ID)
		resp := &sarama.ProducerMessage{
			Topic: "pong",
			Key:   sarama.StringEncoder(receivedMessage.ID),
			Value: sarama.StringEncoder(responseText),
		}
		_, _, err = producer.SendMessage(resp)
		if err != nil {
			log.Printf("Failed to send message to Kafka: %v", err)
		}
	}
	log.Println("Channel closed, exiting")
}
