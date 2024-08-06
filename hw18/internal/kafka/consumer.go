package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"main/hw18/internal/models"
)

func StartConsumer() {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "oranges",
		GroupID: "orange-consumers",
	})

	defer reader.Close()

	basket := models.Basket{}
	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for range ticker.C {
			fmt.Printf("Oranges: small=%d, medium=%d, large=%d\n", basket.Small, basket.Medium, basket.Large)
		}
	}()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Failed to read message from Kafka: %s", err)
			continue
		}

		var orange models.Orange
		if err := json.Unmarshal(msg.Value, &orange); err != nil {
			log.Printf("Failed to unmarshal orange: %s", err)
			continue
		}

		switch {
		case orange.Size < 5:
			basket.Small++
		case orange.Size >= 5 && orange.Size <= 8:
			basket.Medium++
		default:
			basket.Large++
		}
	}
}
