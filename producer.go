package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to local NATS
	nc, err := nats.Connect("nats://localhost:4222,nats://localhost:4223")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Create JetStream cotext
	js, err := nc.JetStream()
	if err != nil {
		log.Fatal(err)
	}

	// Stream name and subject name
	streamName := "MESSAGES"
	subject := "messages"

	_, err = js.AddStream(&nats.StreamConfig{
		Name: streamName,
		Subjects: []string{subject},
		Storage: nats.FileStorage,
	})
	
	if err != nil && err != nats.ErrStreamNameAlreadyInUse {
		log.Fatal(err)
	}

	// Publish a message
	_, err = js.Publish(subject, []byte("Hello from Capventis!"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Message published successfully!")
}