package main

import (
	"learn/redpanda/exercise1"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
    wg.Add(2) // We have two goroutines

	go func() {
        defer wg.Done() // Signal that this goroutine is done
        exercise1.Produce() // Start the producer
    }()

    go func() {
        defer wg.Done() // Signal that this goroutine is done
        exercise1.Consume() // Start the consumer
    }()

    wg.Wait() // Wait for both goroutines to finish
    log.Println("Main function completed.")
}