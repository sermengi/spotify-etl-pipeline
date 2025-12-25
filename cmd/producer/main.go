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

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker},
		Topic: "spotify-events",
		Balancer: &kafka.LeastBytes{},
	})
	defer w.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := w.WriteMessages(ctx,
		kafka.Message{
			Key: []byte("episode:123"),
			Value: []byte(`{"type":"episode_fetched","episode_id":"123"}`),
		},
	)
	if err != nil{
		log.Fatalf("failed to write message: %v", err)
	}

	log.Println("message successfully produced")
}