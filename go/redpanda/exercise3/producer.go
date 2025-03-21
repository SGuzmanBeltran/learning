package exercise3

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/segmentio/kafka-go"
)

// CustomPartitioner implements the Balancer interface
type CustomPartitioner struct{}

func (p CustomPartitioner) Balance(msg kafka.Message, partitions ...int) int {
	if partition, ok := StockPartitions[string(msg.Key)]; ok {
		return partition
	}
	// Default to the first partition if the key is not in our map
	return partitions[0]
}

func Produce(partitionName string, stocks []Stock, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create a new writer (producer)
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{Config.BrokerAddress},
		Topic:    Config.TopicName,
		Balancer: CustomPartitioner{}, // You can choose different balance strategies
	})

	// Ensure the writer is closed after use
	defer func() {
		if err := writer.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
	}()

	log.Printf("Producer for %s - Connected \n", partitionName)

	for i, stock := range stocks {
		fStock := fmt.Sprintf("[%d - %d - %t]", stock.Value, stock.Timestamp, stock.Filter)
		err := writer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(partitionName),
				Value: []byte(fStock),
			},
		)

		if err != nil {
			log.Fatalf("Procuder for %s - failed to write message %d: %v", partitionName, i, err)
		}

		log.Printf("Producer for %s - Produce: %d\n", partitionName, stock.Value)
	}

	log.Printf("All messages for %s produced.\n", partitionName)
}
