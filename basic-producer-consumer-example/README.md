# Basic Producer-Consumer Example with Kafka

This repository demonstrates a basic producer-consumer pattern using Apache Kafka in Go. The example simulates an order creation event where the **producer** sends an "Order Created" event to a Kafka topic, and the **consumer** listens for these events to process them (e.g., sending notifications).

## Project Structure

```
basic-producer-consumer-example/
├── consumer/
│   ├── main.go       # Consumer application that listens to Kafka events
│   └── go.mod        # Go module file for the consumer
├── producer/
│   ├── main.go       # Producer application that sends events to Kafka
│   └── go.mod        # Go module file for the producer
└── README.md         # This file
```

---

## Prerequisites

Before running the project, ensure you have the following installed:

1. **Go**: Install Go from [https://golang.org/doc/install](https://golang.org/doc/install).
2. **Kafka**: Install and run Apache Kafka locally or use a Docker container.
   - Follow the official Kafka quickstart guide: [https://kafka.apache.org/quickstart](https://kafka.apache.org/quickstart).
3. **Zookeeper**: Kafka requires Zookeeper to be running.

---

## Running the Application

### 1. Build and Run the Producer

Navigate to the `producer` directory and run the producer:

```bash
cd producer
go mod tidy
go run main.go
```

The producer will send an "Order Created" event to the Kafka topic `orders`. You should see output similar to:

```
Событие 'Order Created' отправлено в Kafka
```

### 2. Build and Run the Consumer

In a separate terminal, navigate to the `consumer` directory and run the consumer:

```bash
cd consumer
go mod tidy
go run main.go
```

The consumer will listen for messages on the `orders` topic. When it receives an event, it will print the order details and simulate sending a notification. You should see output similar to:

```
Ожидание событий 'Order Created'...
Получено событие 'Order Created': {ID:order-123 UserID:user-456 ProductID:product-789 Amount:100.5 CreatedAt:2023-10-01 12:34:56.789}
Отправка уведомления пользователю user-456 о заказе order-123
```

---

## Code Explanation

### Producer (`producer/main.go`)

- The producer creates an `Order` object and serializes it into JSON format.
- It sends the serialized order as a message to the Kafka topic `orders`.
- The `key` of the message is set to the `UserID`, which can be used for partitioning.

### Consumer (`consumer/main.go`)

- The consumer subscribes to the Kafka topic `orders` using a consumer group (`notification-service`).
- It reads messages from the topic, deserializes the JSON payload into an `Order` object, and processes it by calling the `sendNotification` function.
- Notifications are simulated by printing a message to the console.

---
