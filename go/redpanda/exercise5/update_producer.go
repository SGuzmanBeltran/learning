package exercise5

import (
	"context"
	"learn/redpanda/common"
	"log"

	"github.com/segmentio/kafka-go"
)

type UpdateProducer struct {
	Config common.ConfigTopic
}

func (up *UpdateProducer) Produce(ride *Ride, status string) {
	writer := kafka.Writer{
		Addr:  kafka.TCP(up.Config.BrokerAddress),
		Topic: up.Config.TopicName,
		Balancer: common.CustomPartitioner{
			Partitioner: Partitioner,
		},
	}

	defer func() {
		if err := writer.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
	}()


	log.Printf("Producer for [%d - %d] - Connected \n", ride.RiderID, ride.DriverID)

	err := writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Update"),
			Value: []byte(status),
		},
	)


	if err != nil {
		log.Fatalf("Procuder for Rider [%d - %d] - failed to write message - %v", ride.RiderID, ride.DriverID, err)
	}

	log.Printf("Producer for [%d - %d] - Connected \n", ride.RiderID, ride.DriverID)
}