package exercise1

import (
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
    wg.Add(2) // We have two goroutines

	go func() {
        defer wg.Done() // Signal that this goroutine is done
        Produce() // Start the prsoducer
    }()

    go func() {
        defer wg.Done() // Signal that this goroutine is done
        Consume() // Start the consumer
    }()

    wg.Wait() // Wait for both goroutines to finish
    log.Println("Main function completed.")
}