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

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "notification-service",
	})

	defer reader.Close()

	fmt.Println("Ожидание событий 'Order Created'...")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Ошибка при чтении сообщения: %v", err)
		}

		var order Order
		err = json.Unmarshal(msg.Value, &order)
		if err != nil {
			log.Printf("Ошибка при десериализации заказа: %v", err)
			continue
		}

		fmt.Printf("Получено событие 'Order Created': %+v\n", order)
		sendNotification(order)
	}
}

func sendNotification(order Order) {
	fmt.Printf("Отправка уведомления пользователю %s о заказе %s\n", order.UserID, order.ID)
}
