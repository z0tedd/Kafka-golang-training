# Kafka Training Repository

This repository has been created for training and learning Apache Kafka in a practical way. It contains several example projects that demonstrate different use cases and patterns for working with Kafka using the Go programming language (`golang`). Each example is designed to help you understand how Kafka can be used in real-world scenarios, while also highlighting potential pitfalls and best practices.

## Repository Structure

The repository is organized into several subdirectories, each containing a specific example project:

1. **basic-listener-example**

   - A simple Kafka listener application that demonstrates how to subscribe to a Kafka topic and process incoming messages.
   - **Key Learning Points**: Basic Kafka consumer setup, message processing, and interaction with Kafka topics.

2. **basic-producer-consumer-example**

   - Demonstrates a basic producer-consumer pattern where one service produces "Order Created" events, and another service consumes these events to simulate sending notifications.
   - **Key Learning Points**: Kafka producer-consumer architecture, event-driven design, and JSON serialization/deserialization.

3. **microservice-simple-example**
   - An example of microservice interaction via Kafka, where an API Gateway sends a "ping" message to a Kafka topic, and a second microservice responds with a "pong" message.
   - **Key Learning Points**: Microservice communication using Kafka, request-response anti-patterns, and architectural considerations for scaling Kafka-based systems.

---

## Prerequisites

Before running any of the examples, ensure you have the following installed:

- **Go** (1.16 or higher)
- **Kafka** (running locally or accessible via network)
- **Zookeeper** (required by Kafka)
- **Docker** and **Docker Compose** (for running Kafka and Zookeeper in containers)

---

## Recommendations for Further Learning

1. **Explore Kafka Features**: Dive deeper into Kafka's features such as partitioning, consumer groups, and exactly-once delivery semantics.
2. **Asynchronous Processing**: Experiment with fully asynchronous message processing patterns to better leverage Kafka's strengths.
3. **Scalability**: Try adding more partitions and consumer groups to see how Kafka scales horizontally.
4. **Real-World Use Cases**: Apply what you've learned to real-world scenarios, such as log aggregation, stream processing, or event sourcing.
