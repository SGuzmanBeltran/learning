package exercise2

import (
	"log"
	"sync"
	"time"
)

type Message struct {
	User string
	Message string
	Timestamp int64
}

func Run() {
	var wg sync.WaitGroup
    wg.Add(3) // We have two goroutines
	topic := "chat-room"

	aliceMessages := []Message{
		{
            User:      "Alice",
            Message:   "Hello, World!",
            Timestamp: 1729125778,
        },
        {
            User:      "Alice",
            Message:   "How's it going?",
            Timestamp: 1729125790,
        },
        {
            User:      "Alice",
            Message:   "Let's meet for lunch.",
            Timestamp: 1729125800,
        },
	}

	bobMessages := []Message{
		{
            User:      "Bob",
            Message:   "Hello, Lady!",
            Timestamp: 1729125785,
        },
        {
            User:      "Bob",
            Message:   "Everything good, what about you?",
            Timestamp: 1729125796,
        },
        {
            User:      "Bob",
            Message:   "Of Course, What about Indian Food???!!!",
            Timestamp: 1729125810,
        },
	}

	go func() {
        defer wg.Done() // Signal that this goroutine is done
        Consume(topic) // Start the consumer
    }()

	time.Sleep(5 * time.Second)


	go func() {
        defer wg.Done() // Signal that this goroutine is done
        Produce(topic, "Alice", aliceMessages) // Start the prsoducer
    }()

	go func() {
        defer wg.Done() // Signal that this goroutine is done
        Produce(topic, "Bob", bobMessages) // Start the prsoducer
    }()


    wg.Wait() // Wait for both goroutines to finish
    log.Println("Main function completed.")
}