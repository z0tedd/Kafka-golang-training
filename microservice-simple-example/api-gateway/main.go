package main

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	_ "gateway/docs" // Import the generated docs package

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/google/uuid"
)

// MyMessage represents the structure of the message being sent/received.
type MyMessage struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

var (
	responseChannels = make(map[string]chan *sarama.ConsumerMessage)
	mu               sync.Mutex
)

// Swagger Metadata
//
//	@title			Ping Service API
//	@version		1.0
//	@description	A simple service that sends a "Ping" message to Kafka and waits for a "Pong" response.
//	@host			localhost:3005
//	@BasePath		/
//	@schemes		http
func main() {
	// Initialize Kafka producer and consumer
	producer, err := sarama.NewSyncProducer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	consumer, err := sarama.NewConsumer([]string{"kafka:9092"}, nil)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}
	defer consumer.Close()

	partConsumer, err := consumer.ConsumePartition("pong", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to consume partition: %v", err)
	}
	defer partConsumer.Close()

	// Goroutine to handle Kafka responses
	go handleKafkaResponses(partConsumer)

	// Initialize Fiber router
	router := fiber.New()

	// Serve Swagger UI
	router.Get("/swagger/*", swagger.HandlerDefault) // Default Swagger UI

	router.Get("/ping", func(c *fiber.Ctx) error {
		return handlePingRequest(c, producer)
	})

	// Start the server
	log.Fatal(router.Listen(":3005"))
}

// handleKafkaResponses listens for messages from Kafka and routes them to the appropriate response channel.
// Define the /ping endpoint
//
//	@Summary		Sends a "Ping" message and waits for a "Pong" response.
//	@Description	This endpoint sends a "Ping" message to a Kafka topic and waits for a corresponding "Pong" response.
//	@Tags			ping
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	map[string]string	"Successful response with the 'Pong' message."
//	@Failure		500	{object}	map[string]string	"Internal server error (e.g., Kafka communication failure or timeout)."
//	@Router			/ping [get]
func handleKafkaResponses(partConsumer sarama.PartitionConsumer) {
	for msg := range partConsumer.Messages() {
		responseID := string(msg.Key)
		mu.Lock()
		ch, exists := responseChannels[responseID]
		if exists {
			ch <- msg
			delete(responseChannels, responseID)
		}
		mu.Unlock()
	}
	log.Println("Kafka response listener exited")
}

// handlePingRequest handles the /ping endpoint logic.
func handlePingRequest(c *fiber.Ctx, producer sarama.SyncProducer) error {
	// Generate a unique request ID
	requestID := uuid.New().String()

	// Create the message payload
	message := MyMessage{
		ID:    requestID,
		Name:  "Ping",
		Value: "Pong",
	}

	// Marshal the message to JSON
	bytes, err := json.Marshal(message)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to marshal JSON"})
	}

	// Send the message to Kafka
	msg := &sarama.ProducerMessage{
		Topic: "ping",
		Key:   sarama.StringEncoder(requestID),
		Value: sarama.ByteEncoder(bytes),
	}
	_, _, err = producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v", err)
		return c.Status(500).JSON(fiber.Map{"error": "failed to send message to Kafka"})
	}

	// Create a response channel and store it in the map
	responseCh := make(chan *sarama.ConsumerMessage)
	mu.Lock()
	responseChannels[requestID] = responseCh
	mu.Unlock()

	// Wait for a response or timeout
	select {
	case responseMsg := <-responseCh:
		return c.Status(200).JSON(fiber.Map{"message": string(responseMsg.Value)})
	case <-time.After(10 * time.Second):
		mu.Lock()
		delete(responseChannels, requestID)
		mu.Unlock()
		return c.Status(500).JSON(fiber.Map{"error": "timeout waiting for response"})
	}
}
