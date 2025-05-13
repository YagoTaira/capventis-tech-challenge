# capventis-tech-challenge

Technical project developed for Capventis as part of the technical assessment process.

This project demonstrates a resilient messaging pipeline using **Go**, **NATS JetStream**, and **PostgreSQL**, orchestrated with **Docker Compose**.  
It includes:

- A **producer** that publishes messages to NATS JetStream
- A **consumer** that receives and stores them in a PostgreSQL database
- A **3-node clustered NATS setup** to ensure failover and high availability

## 🚀 Technologies Used

- [Go](https://go.dev/) (v1.24.2)
- [NATS JetStream](https://docs.nats.io/nats-concepts/jetstream)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker Compose](https://docs.docker.com/compose/)
- [godotenv](https://github.com/joho/godotenv) (for environment variables)

## 📦 Project Structure

```bash
capventis-tech-challenge/
├── docker-compose.yml      # NATS cluster + Postgres setup
├── producer.go             # Go app to publish messages to NATS JetStream
├── consumer.go             # Go app to consume and store messages
├── go.mod / go.sum         # Go module dependencies
├── .env                    # DB credentials (gitignored)
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
DATABASE_URL=postgres://user:password@localhost:5432/messages_db
```

### 3. Start NATS + Postgres with Docker Compose

```bash
docker compose up
```

This will start:

- A 3-node NATS JetStream cluster (nats1, nats2, nats3)
- A Postgres database (localhost:5432)

### 4. Run the Consumer

In a new terminal (with Go installed):

```bash
go run consumer.go
```

This connects to JetStream and PostgreSQL, listening to messages and saving them to the database.

### 5. Run the Producer

In another terminal window:

```bash
go run producer.go
```

It will send a “Hello from Capventis!” message to JetStream, which the consumer will receive and persist.

## 🛡️ Resilience Testing (Coming Soon)

The NATS cluster includes 3 nodes to support failover. You can test resilience like this:

1. Start the full system (docker compose up)
2. Run consumer.go and producer.go
3. Stop a node:
   ```bash
   docker stop nats1
   ```
4. Continue publishing messages:
   ```bash
   go run producer.go
   ```
   ✅ Messages will still be delivered and saved in the database.

## 📝 Notes

- The Go consumer auto-creates the messages table if it doesn't exist.
- Message acknowledgments are only sent to JetStream after a successful DB insert.
- Environment variables (via .env) are loaded using godotenv.
- PostgreSQL data is persisted via a Docker volume
- You can query saved messages using psql:

  ```bash
  psql -h localhost -U user -d messages_db
  SELECT * FROM messages;
  ```

## 📚 Useful Resources

- [NATS Documentation](https://docs.nats.io)
- [JetStream Concepts](https://docs.nats.io/nats-concepts/jetstream)
- [Go Programming Language](https://go.dev/doc/)
- [Docker Documentation](https://docs.docker.com)
- [Postgres Documentation](https://www.postgresql.org/docs/)
- [pgx PostgreSQL Driver](https://github.com/jackc/pgx)
- [dotenv](https://github.com/joho/godotenv)

## 📩 Contact

For any questions regarding this project, please contact:

Yago Taira

[GitHub Profile](https://github.com/YagoTaira)
