package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var w *kafka.Writer

// NewProducer 初始化 Kafka Writer
func NewProducer(brokers []string, topic string) {
	w = kafka.NewWriter(kafka.WriterConfig{
		Brokers:  brokers,
		Topic:    topic,
		Balancer: &kafka.LeastBytes{}, // 負載均衡策略
	})
}

// GetWriter 取得全域 Writer
func GetWriter() *kafka.Writer {
	return w
}

// ProduceMessage 發送事件
func ProduceMessage(ctx context.Context, key string, value []byte) error {
	if w == nil {
		log.Fatal("kafka writer not initialized")
		return nil
	}

	msg := kafka.Message{
		Key:   []byte(key),
		Value: value,
		Time:  time.Now(),
	}

	return w.WriteMessages(ctx, msg)
}

// CloseProducer 關閉 Writer（graceful shutdown）
func CloseProducer() {
	if w != nil {
		if err := w.Close(); err != nil {
			log.Println("Error closing Kafka writer:", err)
		}
	}
}

//進入容器
//docker exec -it kafka /bin/bash
//查看topic
// /opt/kafka/bin/kafka-topics.sh --bootstrap-server localhost:9092 --list
//查看groups
// /opt/kafka/bin/kafka-consumer-groups.sh --bootstrap-server localhost:9092 --list
//查看broker
// /opt/kafka/bin/kafka-broker-api-versions.sh --bootstrap-server kafka:9092
