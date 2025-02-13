# Basic Kafka Listener Example

This is a simple Go application that listens to a Kafka topic and prints the messages it receives. It uses the `segmentio/kafka-go` library to interact with Kafka.

## Prerequisites

Before running this example, ensure you have the following installed:

- **Go** (1.16 or higher)
- **Kafka** (running locally or accessible via network)
- **kafka-go** library (`github.com/segmentio/kafka-go`)

You can install the required Go package using:

```bash
go get github.com/segmentio/kafka-go
```

## Running the example

```bash
go run main.go
```
