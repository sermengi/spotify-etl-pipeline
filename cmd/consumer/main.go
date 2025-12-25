package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	broker := os.Getenv("KAFKA_BROKER")
	if broker == "" {
		broker = "localhost:9092"
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic: "spotify-events",
		GroupID: "spotify-etl-local",
	})
	defer r.Close()

	log.Println("consumer started")

	for {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		msg, err := r.ReadMessage(ctx)
		cancel()

		if err != nil {
			log.Printf("read error: %v", err)
			continue
		}

		log.Printf(
			"received message topic=%s partition=%d offset=%d key=%s value=%s",
			msg.Topic,
			msg.Partition,
			msg.Offset,
			string(msg.Key),
			string(msg.Value),
		)
	}
}