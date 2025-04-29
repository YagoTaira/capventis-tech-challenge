package main

import (
	"log"

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

	// Subscribe to the 'messages' subject
	sub, err := js.Subscribe("messages", func(msg *nats.Msg) {
		log.Printf("Received message: %s\n", string(msg.Data))
		msg.Ack()
	}, nats.Durable("my-consumer"), nats.ManualAck())
	if err != nil {
		log.Fatal(err)
	}

	defer sub.Unsubscribe()

	log.Println("Listening for messages... (press Ctrl+C to exit)")
	select {} // block forever
}