package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"order-service-platform/kafka"
	"order-service-platform/proto/proto/pb"
	"order-service-platform/worker/order/model"
	"time"

	"github.com/google/uuid"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
}

func (s *OrderServer) Order(c context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	fmt.Printf("%d", req.EventId)
	msg := model.OrderMessage{
		OrderID:   uuid.NewString(), // server 產生
		EventID:   int(req.EventId),
		UserID:    req.UserId,
		Quantity:  int(req.Quantity),
		Timestamp: time.Now().Unix(),
	}

	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println("marshal order message failed:", err)
		return nil, err
	}

	err = kafka.ProduceMessage(c, kafka.TopicOrder, payload)
	if err != nil {
		log.Println("Kafka publish error:", err, " req:", req.String())
		return nil, err
	}

	log.Println("Kafka publish success. req:", req.String())

	return &pb.OrderResponse{}, nil
}
