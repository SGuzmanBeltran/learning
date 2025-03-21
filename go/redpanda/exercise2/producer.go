package exercise2

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func Produce(topic string, user string, messages []Message) {
	// Kafka broker address and topic
	brokerAddress := "localhost:9092" // Update with your Redpanda broker address

	// Create a new writer (producer)
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{brokerAddress},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{}, // You can choose different balance strategies
	})

	// Ensure the writer is closed after use
	defer func() {
		if err := writer.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
	}()

	log.Println("Producer - Connected")

	for i, message := range messages {
		fMessage := fmt.Sprintf("[%s - %d]: %s", message.User, message.Timestamp, message.Message)
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(fmt.Sprintf("[%s - %d - %d]", message.User, message.Timestamp, i)),
				Value: []byte(fMessage),
			},
		)

		if err != nil {
			log.Fatalf("failed to write message %d: %v", i, err)
		}

		log.Printf("Producer - Produce: %s - %d\n", message.User, i+1)
	}

	log.Println("All messages produced.")
}
