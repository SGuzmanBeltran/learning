package exercise5

import (
	"context"
	"learn/redpanda/common"
	"log"
	"sync"

	"github.com/segmentio/kafka-go"
)

type RiderProducer struct {
	Config common.ConfigTopic
}

func (rp *RiderProducer) Produce(rider Rider, wg *sync.WaitGroup) {
	defer wg.Done()

	writer := kafka.Writer{
		Addr:  kafka.TCP(rp.Config.BrokerAddress),
		Topic: rp.Config.TopicName,
		Balancer: common.CustomPartitioner{
			Partitioner: Partitioner,
		},
	}

	defer func() {
		if err := writer.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
	}()

	log.Printf("Producer for %d - Connected \n", rider.ID)

	value, err := common.ToJSON(rider)
	if err != nil {
		log.Fatalf("Producer for %d - failed to pass rider to json", rider.ID)
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte("Rider"),
			Value: value,
		},
	)

	if err != nil {
		log.Fatalf("Procuder for Rider %d - failed to write message - %v", rider.ID, err)
	}

	log.Printf("Producer for Rider %d - Produce \n", rider.ID)
}
