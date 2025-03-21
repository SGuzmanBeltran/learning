package preexercises

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func Consume() {
	brokerAddress := "localhost:9092" // Update with your Redpanda broker address
	topic := "test"
	groupID := "my-consumer-group"

	// Create a new reader with the topic and the broker address
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{brokerAddress},
		Topic:       topic,
		GroupID:  groupID, // Set the GroupID here
		MinBytes:    10e3, // 10KB
		MaxBytes:    1e6,  // 1MB
	})

	log.Println("Consumer - Connected")

	defer reader.Close()

	for {
		// Read a message from the topic
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("failed to read message: %v", err)
			break
		}
		fmt.Printf("Consumer - Received: %s\n", string(m.Value))

		// Commit the offset after processing the message
		if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Printf("failed to commit message: %v", err)
		}
	}
}
