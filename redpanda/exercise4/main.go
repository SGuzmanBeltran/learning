package exercise4

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

type Stock struct {
	Value     int64
	Timestamp int64
	Filter    bool
}

type ConfigTopic struct {
	TopicName     string
	Partitions    int64
	BrokerAddress string
}

var StockPartitions = map[string]int{
	"AAPL":  0,
	"GOOGL": 1,
	"TSLA":  2,
}

var Config = ConfigTopic{
	BrokerAddress: "localhost:9092",
	TopicName:     "stock-prices",
	Partitions:    3,
}

func Run() {
	setTopic()

	googleStocks := []Stock{
		{
			Value:     150,
			Timestamp: 1729125785,
			Filter:    true,
		},
		{
			Value:     140,
			Timestamp: 1729125796,
			Filter:    false,
		},
		{
			Value:     130,
			Timestamp: 1729125810,
			Filter:    false,
		},
		{
			Value:     1200,
			Timestamp: 1729125920,
			Filter:    true,
		},
		{
			Value:     1000,
			Timestamp: 1729125940,
			Filter:    true,
		},
	}

	appleStocks := []Stock{
		{
			Value:     300,
			Timestamp: 1729125785,
			Filter:    true,
		},
		{
			Value:     310,
			Timestamp: 1729125796,
			Filter:    true,
		},
		{
			Value:     200,
			Timestamp: 1729125810,
			Filter:    true,
		},
		{
			Value:     500,
			Timestamp: 1729125920,
			Filter:    true,
		},
		{
			Value:     700,
			Timestamp: 1729125940,
			Filter:    true,
		},
	}

	teslaStocks := []Stock{
		{
			Value:     5000,
			Timestamp: 1729125785,
			Filter:    true,
		},
		{
			Value:     300,
			Timestamp: 1729125796,
			Filter:    false,
		},
		{
			Value:     100,
			Timestamp: 1729125810,
			Filter:    true,
		},

		{
			Value:     120,
			Timestamp: 1729125920,
			Filter:    true,
		},

		{
			Value:     160,
			Timestamp: 1729125940,
			Filter:    true,
		},
	}

	var wg sync.WaitGroup

	// Start consumers
	wg.Add(3)                // 3 consumers
	go Consume("AAPL", &wg)  // Start the consumer
	go Consume("GOOGL", &wg) // Start the consumer
	go Consume("TSLA", &wg)  // Start the consumer

	time.Sleep(5 * time.Second)

	// Start producers
	wg.Add(3)                              // 3 producers
	go Produce("AAPL", appleStocks, &wg)   // Start the prsoducer
	go Produce("GOOGL", googleStocks, &wg) // Start the prsoducer
	go Produce("TSLA", teslaStocks, &wg)   // Start the prsoducer

	// Wait for all goroutines to finish
	wg.Wait()

	log.Println("Main function completed.")
}

func setTopic() {
	// Create a new Kafka client
	conn, err := kafka.Dial("tcp", Config.BrokerAddress)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	defer conn.Close()

	// Create a new topic configuration
	topicConfig := kafka.TopicConfig{
		Topic:             Config.TopicName,
		NumPartitions:     int(Config.Partitions),
		ReplicationFactor: 1,
	}

	// Create the topic
	err = conn.CreateTopics(topicConfig)
	if err != nil {
		log.Fatal("failed to create topic:", err)
	}

	fmt.Println("Topic created successfully!")
}
