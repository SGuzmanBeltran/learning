package preexercises

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func Produce() {
	// Kafka broker address and topic
	brokerAddress := "localhost:9092" // Update with your Redpanda broker address
	topic := "test"

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

	for i := 1; i <= 5; i++ {
		message := fmt.Sprintf("Message %d", i)
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(fmt.Sprintf("Key-%d", i)),
				Value: []byte(message),
			},
		)

		if err != nil {
			log.Fatalf("failed to write message %d: %v", i, err)
		}

		log.Printf("Produced: %s\n", message)

		// Optional: Add a delay between messages
		time.Sleep(1 * time.Second)
	}

	log.Println("All messages produced.")
}
