package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Order struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	ProductID string    `json:"product_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	brokerAddress := "localhost:9092"
	topic := "orders"

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
	})

	defer writer.Close()

	order := Order{
		ID:        "order-123",
		UserID:    "user-456",
		ProductID: "product-789",
		Amount:    100.50,
		CreatedAt: time.Now(),
	}

	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Fatalf("Ошибка при сериализации заказа: %v", err)
	}

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(order.UserID),
		Value: orderJSON,
	})
	if err != nil {
		log.Fatalf("Ошибка при отправке сообщения: %v", err)
	}

	fmt.Println("Событие 'Order Created' отправлено в Kafka")
}
