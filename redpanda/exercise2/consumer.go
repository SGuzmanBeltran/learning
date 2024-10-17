package exercise2

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/segmentio/kafka-go"
)

func Consume(topic string) {
	brokerAddress := "localhost:9092" // Update with your Redpanda broker address
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

	messages := []Message{}

	for {
		// Read a message from the topic
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("failed to read message: %v", err)
			break
		}
		log.Printf("Consumer - Received: %s\n", string(m.Value))
		message := getMessage(string(m.Value))
		messages = append(messages, message)


		// Commit the offset after processing the message
		if err := reader.CommitMessages(context.Background(), m); err != nil {
			log.Printf("failed to commit message: %v", err)
		}
	}


	//Process at in timestamp
}

func getMessage(inputMessage string) (Message) {
	parts := strings.SplitN(strings.Trim(inputMessage, "[]"), "]: ", 2)
    metadata := parts[0]
    message := parts[1]

    // Step 2: Split the metadata part to get user and timestamp
    metaParts := strings.SplitN(metadata, " - ", 2)
    user := metaParts[0]
    timestampStr := metaParts[1]

    // Step 3: Convert timestamp to int64
    timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
    if err != nil {
        fmt.Println("Error parsing timestamp:", err)
        return Message{}
	}

	return Message{
		User: user,
		Message: message,
		Timestamp: timestamp,
	}
}
