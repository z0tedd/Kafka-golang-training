# Project "Ping-Pong via Kafka"

## Project Description

This project is an example of implementing interaction between two microservices using Apache Kafka. In this case, one microservice (API Gateway) sends a message to the `ping` topic, and the second microservice (`second-microservice`) reads that message, processes it, and sends a response to the `pong` topic. The API Gateway waits for a response from the `pong` topic and returns it to the client.

### Project Architecture

1. **Zookeeper** — Used for managing the Kafka cluster.
2. **Kafka** — A message broker that handles messages between microservices.
3. **API Gateway** — An HTTP server that accepts client requests, sends them to Kafka, and waits for a response.
4. **Second Microservice** — A service that reads messages from the `ping` topic, processes them, and sends a response to the `pong` topic.

---

## Running the Project

To run the project, follow these steps:

1. Ensure that Docker and Docker Compose are installed on your machine.
2. Clone the project repository.
3. Run the following command:

   ```bash
   docker-compose up
   ```

4. After the containers have successfully started, the API Gateway will be available at `http://localhost:3005`.

---

## Usage

### Swagger UI

Swagger UI is available at:

```
http://localhost:3005/swagger/index.html
```

### Example Request

Send a GET request to the `/ping` endpoint:

```bash
curl http://localhost:3005/ping
```

You will receive a JSON response, for example:

```json
{
  "message": "Ping Pong (some-uuid)"
}
```

---

## Implementation Details

### wait-for-it.sh

The project uses the `wait-for-it.sh` script, which allows microservices to wait until Kafka is fully ready before starting. This is important because the microservices depend on Kafka for their operation.

The original script code can be found in the repository:  
[https://github.com/vishnubob/wait-for-it](https://github.com/vishnubob/wait-for-it)

Example usage in the Dockerfile:

```dockerfile
CMD ["./wait-for-it.sh", "kafka:9092", "--", "./main"]
```

---

## Problems and Shortcomings of the Project

This project resembles a collection of anti-patterns in microservice development. Here's why:

1. **API Gateway Cannot Be Scaled**

   - The API Gateway stores state (response channels) in memory. If it crashes or restarts, all pending requests will be lost. This makes it unreliable and unsuitable for scaling.

2. **Improper Use of Kafka**

   - In Kafka, keys (the `key` field) are used to select partitions. However, in this project, only one partition (`partition 0`) is used, making it impossible to scale Kafka.
   - Consumer groups should be used for message processing, but this project does not implement them.

3. **Inability to Scale the Second Microservice**

   - The second microservice (`second-microservice`) is tied to a single partition (`partition 0`). This means that even if you run multiple instances of the microservice, they won't be able to process messages in parallel because each instance will compete for the same data.

4. **Request-Response Pattern via Broker — Bad Practice**
   - Using Kafka to implement the "request-response" pattern is considered an anti-pattern. Message brokers are designed for asynchronous data processing, where exact delivery and processing times are not critical.
   - If you need a synchronous response to a request, it's better to use a load balancer (e.g., NGINX) and direct HTTP requests to services.

---

## Recommendations for Improvement

1. **Use HTTP for Synchronous Requests**  
   If you need to get a response to a request, use the HTTP protocol and a load balancer. This will help avoid the complexities associated with using Kafka for synchronous communication.

2. **Scale Kafka**

   - Add more partitions to the topics to distribute the load.
   - Use consumer groups for parallel message processing.

3. **State in API Gateway**

   - Avoid storing state in the API Gateway's memory. Instead, use an external storage system (e.g., Redis) to store response channels.

4. **Asynchronous Processing**
   - Reconsider the project architecture. If possible, switch to a fully asynchronous message processing model, where Kafka is used as intended — for "eventual" message delivery.

---

## Conclusion

This project demonstrates basic microservice interaction via Kafka but contains numerous architectural flaws. It can be useful for learning and understanding how Kafka works, but it is not recommended for use in real projects without significant improvements.
