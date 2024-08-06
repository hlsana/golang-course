package kafka

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
	"main/hw18/internal/models"
)

func StartProducer() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "oranges",
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()

	for {
		orange := models.Orange{
			Size: rand.Float64()*10 + 2, 
		}

		orangeBytes, err := json.Marshal(orange)
		if err != nil {
			log.Printf("Failed to marshal orange: %s", err)
			continue
		}

		err = writer.WriteMessages(context.Background(),
			kafka.Message{
				Value: orangeBytes,
			},
		)
		if err != nil {
			log.Printf("Failed to write message to Kafka: %s", err)
			continue
		}

		log.Printf("Produced orange with size: %.2f cm", orange.Size)

		time.Sleep(1 * time.Second)
	}
}
