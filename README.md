# capventis-tech-challenge

Technical project developed for Capventis as part of the technical assessment process.

This project demonstrates a simple messaging pipeline using **Go**, **NATS JetStream**, **PostgreSQL**, and **Docker Compose**. It includes a **producer** that publishes data to JetStream and a **consumer** that listens for messages and stores them in a Postgres database.

## 🚀 Technologies Used

- [Go](https://go.dev/) (v1.24.2)
- [NATS JetStream](https://docs.nats.io/nats-concepts/jetstream)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker Compose](https://docs.docker.com/compose/)

## 📦 Project Structure

```bash
capventis-tech-challenge/
├── docker-compose.yml      # Docker setup for NATS + Postgres
├── producer.go             # Go program to send messages to NATS JetStream
├── consumer.go             # Go program to receive messages and store in DB
├── go.mod                  # Go module configuration
├── .env                    # (gitignored) stores DB credentials
├── README.md               # Project documentation
```

## ⚙️ How to Run the Project

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/capventis-tech-challenge.git
cd capventis-tech-challenge
```

### 2. Create a .env file

Create a .env file in the root folder (this is not tracked in Git):

```bash
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_DB=messages_db
```

### 3. Start NATS + Postgres with Docker Compose

```bash
docker compose up
```

This will start:

- A NATS server with JetStream enabled (localhost:4222)
- A Postgres database (localhost:5432)

### 4. Run the Producer

In a new terminal window:

```bash
go run producer.go
```

This sends a sample message to the JetStream messages stream.

### 5. Run the Consumer

In another terminal window:

```bash
DATABASE_URL=postgres://user:password@localhost:5432/messages_db go run consumer.go
```

This listens for messages from JetStream and stores them into the messages table in Postgres.

## 🛡️ Resilience Testing (Coming Soon)

The project will be expanded to a NATS cluster with multiple nodes to simulate resilience (e.g., shutting down a node without losing messages).

## 📝 Notes

- The Go consumer auto-creates the messages table if it doesn't exist.
- Message acknowledgments are only sent to JetStream after a successful DB insert.
- Database credentials are injected via environment variables using a .env file (excluded from Git).
- You can query saved messages using psql:

```bash
psql -h localhost -U user -d messages_db
SELECT * FROM messages;
```

## 📚 Useful Resources

- [NATS Documentation](https://docs.nats.io)
- [Go Programming Language](https://go.dev/doc/)
- [Docker Documentation](https://docs.docker.com)
- [pgx Postgres Driver](https://github.com/jackc/pgx)

## 📩 Contact

For any questions regarding this project, please contact:

Yago Taira

[GitHub Profile](https://github.com/YagoTaira)
