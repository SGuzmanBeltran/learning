package exercise4

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	movingaverage "github.com/RobinUS2/golang-moving-average"
	"github.com/segmentio/kafka-go"
)

func lastElements(slice []Stock, n int) []Stock {
	length := len(slice)
	if length == 0 || n <= 0 {
		return []Stock{} // return empty slice if original slice is empty or n is non-positive
	}
	if n >= length {
		return slice // if n is greater than or equal to length, return the entire slice
	}
	return slice[length-n:] // return the last n elements
}

func calculateMovingAverage(slice []Stock) float64 {
	ma := movingaverage.New(2)
	if len(slice) < 3 {
		return 0
	}
	for _, v := range slice {
		ma.Add(float64(v.Value))
	}
	return ma.Avg()
}

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

	stocks := []Stock{}

	for {
		// Read a message from the topic
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("failed to read message: %v", err)
			break
		}
		stock := getStock(string(m.Value))
		log.Printf("Consumer for %d - Received: %s, %d\n", partition, string(m.Key), stock.Value)
		stocks = append(stocks, stock)
		lastPrices := lastElements(stocks, 3)
		movingAverage := calculateMovingAverage(lastPrices)
		log.Printf("Consumer for %d - Processing: %s - %f \n", partition, string(m.Key), movingAverage)

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
