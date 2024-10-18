package exercise5

import (
	"context"
	"fmt"
	"learn/redpanda/common"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

type DriverProducer struct {
	Config common.ConfigTopic
}

func (dp *DriverProducer) Produce(driver Driver, wg *sync.WaitGroup) {
	defer wg.Done()

	//Use this writer
	writer := &kafka.Writer{
		Addr:  kafka.TCP(dp.Config.BrokerAddress),
		Topic: dp.Config.TopicName,
	}

	defer func() {
		if err := writer.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
	}()

	log.Printf("Producer for %d - Connected \n", driver.ID)

	if driver.Available {
		return
	}

	fDriver := fmt.Sprintf("[%d]", driver.ID)
	value, err := common.ToJSON(driver)
	if err != nil {
		log.Fatalf("Producer for %d - failed to pass driver to json", driver.ID)
	}

	err = writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(fDriver),
			Value: value,
		},
	)

	if err != nil {
		log.Fatalf("Procuder for %d - failed to write message %v", driver.ID, err)
	}

	log.Printf("Producer for %d - Produce\n", driver.ID)

	time.Sleep(1 * time.Second)
}
