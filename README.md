# capventis-tech-challenge

Technical project developed for Capventis as part of the technical assessment process.

This project demonstrates a simple messaging pipeline using Go, NATS JetStream, and Docker Compose. It includes a producer that publishes data to JetStream and a consumer that listens for messages.

## 🚀 Technologies Used

- [Go](https://go.dev/) (v1.24.2)
- [NATS JetStream](https://docs.nats.io/nats-concepts/jetstream)
- [Docker Compose](https://docs.docker.com/compose/)

## 📦 Project Structure

```bash
capventis-tech-challenge/
├── docker-compose.yml      # Docker setup for NATS server
├── producer.go             # Go program to send messages to NATS JetStream
├── consumer.go             # Go program to consume messages from NATS JetStream
├── go.mod                  # Go module configuration
├── README.md               # Project documentation
```

## ⚙️ How to Run the Project

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/capventis-tech-challenge.git
cd capventis-tech-challenge
```

### 3. Start NATS Server with Docker Compose

```bash
docker compose up
```

This will start a NATS server with JetStream enabled and available on localhost:4222.

### 5. Run the Producer

In a new terminal window:

```bash
go run producer.go
```

This sends a sample message to the messages stream on JetStream.

### 4. Run the Consumer

In another terminal window:

```bash
go run consumer.go
```

The consumer listens for messages and prints them to the terminal.

## 🛡️ Resilience Testing (Coming Soon)

The project will be expanded to a NATS cluster with multiple nodes to simulate resilience (e.g., shutting down a node without losing messages).

## 📝 Notes

- The producer and consumer both connect to the NATS server at localhost:4222.
- Manual acknowledgments are used to ensure reliable delivery of messages.

## 📚 Useful Resources

- [NATS Documentation](https://docs.nats.io)
- [Go Programming Language](https://go.dev/doc/)
- [Docker Documentation](https://docs.docker.com)

## 📩 Contact

For any questions regarding this project, please contact:

Yago Taira

[GitHub Profile](https://github.com/YagoTaira)
