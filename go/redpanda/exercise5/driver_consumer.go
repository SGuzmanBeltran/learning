package exercise5

import (
	"context"
	"encoding/json"
	"learn/redpanda/common"
	"log"
	"sync"

	"github.com/segmentio/kafka-go"
)

type DriverConsumer struct {
	Config common.ConfigTopic
}

func (dc *DriverConsumer) Consume(wg *sync.WaitGroup) {
	defer wg.Done()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{dc.Config.BrokerAddress},
		Topic:    dc.Config.TopicName,
		MinBytes: 10e3, // 10KB
		MaxBytes: 1e6,  // 1MB
	})
	defer reader.Close()

	log.Printf("Consumer for DRIVER - Connected \n")

	for {
		// Read a message from the topic
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("failed to read message: %v", err)
			break
		}

		// Unmarshal the JSON message into the Driver struct
		var driver *Driver
		if err := json.Unmarshal(m.Value, &driver); err != nil {
			log.Printf("Error unmarshalling JSON: %v", err)
			continue
		}
		log.Printf("Consumer for DRIVER - Received: %s\n", string(m.Key))

		Central.ChangeDriverAvailability(driver)
	}
}
