package common

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type ConfigTopic struct {
	TopicName     string
	BrokerAddress string
	Partitions    int64
}

func ToJSON(v interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal struct to JSON: %w", err)
	}
	return jsonData, nil
}

func SetTopic(ct ConfigTopic) {
	// Create a new Kafka client
	conn, err := kafka.Dial("tcp", ct.BrokerAddress)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	defer conn.Close()

	// Create a new topic configuration
	topicConfig := kafka.TopicConfig{
		Topic:         ct.TopicName,
		NumPartitions: int(ct.Partitions),
	}

	// Create the topic
	err = conn.CreateTopics(topicConfig)
	if err != nil {
		log.Fatal("failed to create topic:", err)
	}

	fmt.Println("Topic created successfully!")
}

// CustomPartitioner implements the Balancer interface
type CustomPartitioner struct{
	Partitioner map[string]int
}

func (p CustomPartitioner) Balance(msg kafka.Message,partitions ...int) int {
	if partition, ok := p.Partitioner[string(msg.Key)]; ok {
		return partition
	}
	// Default to the first partition if the key is not in our map
	return partitions[0]
}