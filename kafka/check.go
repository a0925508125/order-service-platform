package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

// WaitUntilReady 確保 Kafka broker + coordinator 可用
func WaitUntilReady(brokers []string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)

	for {
		conn, err := kafka.Dial("tcp", brokers[0])
		if err == nil {
			_ = conn.Close()
			log.Println("Kafka broker is reachable")
			return nil
		}

		if time.Now().After(deadline) {
			return err
		}

		log.Printf("Waiting for Kafka %s to be ready...",brokers[0])
		time.Sleep(2 * time.Second)
	}
}

func WaitUntilGroupCoordinatorReady(
	brokers []string,
	topic string,
	groupID string,
	timeout time.Duration,
) error {

	deadline := time.Now().Add(timeout)

	for {
		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers: brokers,
			Topic:   topic,
			GroupID: groupID,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		_, err := r.ReadMessage(ctx)
		cancel()
		_ = r.Close()

		if err == nil || err == context.DeadlineExceeded {
			log.Println("Kafka group coordinator is ready")
			return nil
		}

		if time.Now().After(deadline) {
			return err
		}

		log.Printf("Waiting for Kafka %s to be ready...",brokers[0])
		time.Sleep(2 * time.Second)
	}
}
