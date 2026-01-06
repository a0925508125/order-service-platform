package kafka

import (
	"fmt"

	"github.com/segmentio/kafka-go"
)

var Brokers = []string{"localhost:9092"} // 可以加多個 broker
const TopicOrder = "topic-order"
const GroupOrder = "group-order"

// createTopic 建立 Topic
func CreateTopic(brokerAddress, topic string, partitions, replicationFactor int) error {
	conn, err := kafka.Dial("tcp", brokerAddress)
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}

	// 連到 controller 才能建立 Topic
	var controllerConn *kafka.Conn
	controllerAddr := fmt.Sprintf("%s:%d", controller.Host, controller.Port)
	controllerConn, err = kafka.Dial("tcp", controllerAddr)
	if err != nil {
		return fmt.Errorf("failed to dial controller: %w", err)
	}
	defer controllerConn.Close()

	// 建立 Topic
	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     partitions,
			ReplicationFactor: replicationFactor,
		},
	}

	return controllerConn.CreateTopics(topicConfigs...)
}
