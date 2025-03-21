package exercise5

import (
	"learn/redpanda/common"
	"sync"
	"time"
)

var ConfigRides = common.ConfigTopic{
	BrokerAddress: "localhost:9092",
	TopicName:     "rides",
	Partitions:    2,
}

var ConfigDrivers = common.ConfigTopic{
	BrokerAddress: "localhost:9092",
	TopicName:     "drivers",
	Partitions:    1,
}

var Partitioner = map[string]int{
	"Rider": 0,
	"Update": 1,
}

func Run() {
	common.SetTopic(ConfigRides)
	common.SetTopic(ConfigDrivers)

	var wg *sync.WaitGroup
	wg.Add(3)
	wg.Add(len(Central.Drivers))
	wg.Add(len(Central.Riders))

	Central.StartCentral(10, 5)

	//Consumers
	driverConsumer := DriverConsumer{
		Config: ConfigDrivers,
	}
	go driverConsumer.Consume(wg)

	riderConsumer := RiderConsumer{
		Config: ConfigRides,
	}
	go riderConsumer.Consume(0, wg)
	go riderConsumer.Consume(1, wg)

	time.Sleep(5 * time.Second)

	//Producers
	driverProducer := DriverProducer{
		Config: ConfigDrivers,
	}
	for _, driver := range Central.Drivers {
		go driverProducer.Produce(driver, wg)
	}

	riderProducer := RiderProducer{
		Config: ConfigDrivers,
	}
	for _, rider := range Central.Riders {
		go riderProducer.Produce(rider, wg)
	}

	wg.Wait()
}
