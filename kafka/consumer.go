package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var r *kafka.Reader

// NewConsumer 建立一個消費 Kafka Reader
func NewConsumer(brokers []string, topic, groupID string) {
	r = kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		GroupID:     groupID,
		MinBytes:    1,
		MaxBytes:    10e6,
		StartOffset: kafka.FirstOffset, // <- 從最早的訊息開始
	})
}

func GetReader() *kafka.Reader {
	return r
}

// ConsumeMessages  消費事件
func ConsumeMessages(ctx context.Context, handleFunc func(context.Context, []byte) error) {
	if r == nil {
		log.Fatal("kafka reader not initialized")
		return
	}
	log.Println("Kafka consumer Messages")

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Println("Kafka consumer stopped")
				return
			default:
				m, err := r.ReadMessage(ctx)
				if err != nil {
					log.Println("Error reading message:", err)
					time.Sleep(1 * time.Second)
					continue
				}

				log.Printf("Received message: key=%s value=%s \n", string(m.Key), string(m.Value))

				// 交給 handler 處理
				if err := handleFunc(ctx, m.Value); err != nil {
					log.Printf("Handler error: %v, message key=%s \n", err, string(m.Key))
				} else {
					log.Printf("Message processed: key=%s \n", string(m.Key))
				}
			}
		}
	}()
}
