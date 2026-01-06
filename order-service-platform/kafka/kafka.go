package kafka

import (
	"log"
)

func InitKafka() {
	// 建立 Topic
	for _, broker := range Brokers {
		if err := CreateTopic(broker, TopicOrder, 1, 10); err != nil {
			log.Fatalf("%s, failed to create Topic: %s, error:%v", broker, TopicOrder, err)
		} else {
			log.Printf("%s, succes to create Topic: %s", broker, TopicOrder)
		}
	}

	NewProducer(Brokers, TopicOrder)
}
