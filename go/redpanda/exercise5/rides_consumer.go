package exercise5

import (
	"context"
	"encoding/json"
	"learn/redpanda/common"
	"log"
	"sync"

	"github.com/segmentio/kafka-go"
)

type RiderConsumer struct {
	Config common.ConfigTopic
}

func (rc *RiderConsumer) Consume(partition int, wg *sync.WaitGroup) {
	defer wg.Done()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{rc.Config.BrokerAddress},
		Topic:    rc.Config.TopicName,
		MinBytes: 10e3, // 10KB
		MaxBytes: 1e6,  // 1MB
	})
	defer reader.Close()

	log.Printf("Consumer for Rider %d - Connected \n", partition)

	for {
		// Read a message from the topic
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("failed to read message: %v", err)
			break
		}

		// Unmarshal the JSON message into the Driver struct
		if string(m.Key) == "Rider" {
			var rider *Rider
			if err := json.Unmarshal(m.Value, &rider); err != nil {
				log.Printf("Error unmarshalling JSON: %v", err)
				continue
			}
			Central.CreateRide(rider)
			ride := Central.MatchRide(rider.ID)
			updateProducer := UpdateProducer{
				Config: rc.Config,
			}
			updateProducer.Produce(ride, "Matched")

		} else if string(m.Key) == "Update" {
			var rideUpdate *RideUpdate
			if err := json.Unmarshal(m.Value, &rideUpdate); err != nil {
				log.Printf("Error unmarshalling JSON: %v", err)
				continue
			}
			Central.UpdateRideState(&rideUpdate.Ride, rideUpdate.Status)
		}

		log.Printf("Consumer for Rider %d - Received: %s\n", partition, string(m.Key))
	}
}
