package exercise3

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

func Consume(partitionName string, wg *sync.WaitGroup) {
	defer wg.Done()
	partition := StockPartitions[partitionName]

	// Create a new reader with the topic and the broker address
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{Config.BrokerAddress},
		Topic:     Config.TopicName,
		MinBytes:  10e3, // 10KB
		MaxBytes:  1e6,  // 1MB
		Partition: partition,
	})

	log.Printf("Consumer for %d - Connected \n", partition)

	defer reader.Close()

	for {
		// Read a message from the topic
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("failed to read message: %v", err)
			break
		}
		stock := getStock(string(m.Value))
		log.Printf("Consumer for %d - Received: %s, %d\n", partition, string(m.Key), stock.Value)
		if !stock.Filter {
			continue
		}

		log.Printf("Consumer for %d - Processing: %s\n", partition, string(m.Key))
		// stocks = append(stocks, stock)

		//Process it, call for update or send notification

		// Commit the offset after processing the message
		// if err := reader.CommitMessages(context.Background(), m); err != nil {
		// 	log.Printf("failed to commit message: %v", err)
		// }

		time.Sleep(1 * time.Second)
	}
}

func getStock(inputStock string) Stock {
	parts := strings.Trim(inputStock, "[]")

	// Step 2: Split the metadata part to get user and timestamp
	metaParts := strings.SplitN(parts, " - ", 3)
	valueStr := metaParts[0]
	timestampStr := metaParts[1]
	filterStr := metaParts[2]

	// Step 3: Convert timestamp to int64
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
	if err != nil {
		fmt.Println("Error parsing timestamp:", err)
		return Stock{}
	}

	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		fmt.Println("Error parsing value:", err)
		return Stock{}
	}

	// Parse the boolean
	filter, err := strconv.ParseBool(filterStr)
	if err != nil {
		fmt.Println("Error parsing boolean:", err)
		return Stock{}
	}

	return Stock{
		Value:     value,
		Timestamp: timestamp,
		Filter:    filter,
	}
}
