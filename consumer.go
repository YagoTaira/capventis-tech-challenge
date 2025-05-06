package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Create JetStream context
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to Postgres
	dbURL := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to Postgres: %v", err)
	}
	defer conn.Close(context.Background())

	// Create table if not exists
	_, err = conn.Exec(context.Background(), `
	CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		content TEXT NOT NULL,
		received_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
	)
	`)
	if err != nil {
		log.Fatalf("Failed to create messages table: %v", err)
	}

	// Subscribe to the 'messages' subject
	sub, err := js.Subscribe("messages", func(msg *nats.Msg) {
		log.Printf("Received message: %s\n", string(msg.Data))

		_, err := conn.Exec(context.Background(),
			"INSERT INTO messages (content) VALUES ($1)",
			string(msg.Data),
		)
		if err != nil {
			log.Printf("Failed to insert messages into DB: %v", err)
			return
		}

		msg.Ack()
		log.Println("Message saves to database.")
	}, nats.Durable("my-consumer"), nats.ManualAck())
	if err != nil {
		log.Fatal(err)
	}

	defer sub.Unsubscribe()

	log.Println("Listening for messages... (press Ctrl+C to exit)")
	select {} // block forever
}