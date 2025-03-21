package exercise1

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
	topic := "fun-messages"

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

	jokes := [5]string{
		"Why did the scarecrow win an award? Because he was outstanding in his field!",
		"Why don’t skeletons fight each other? They don’t have the guts!",
		"What do you call fake spaghetti? An impasta!",
		"Why did the bicycle fall over? Because it was two-tired!",
		"What did one wall say to the other wall? I’ll meet you at the corner!",
	}

	log.Println("Producer - Connected")

	time.Sleep(5 * time.Second) // Adjust as necessary
	// Send messages
	for i, joke := range jokes {
		message := fmt.Sprintf("Message - %s", joke)
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
