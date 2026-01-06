package main

import (
	"fmt"
	"log"
	"net"
	"order-service-platform/kafka"
	"order-service-platform/proto/proto/pb"
	"order-service-platform/service/order/handler"

	"google.golang.org/grpc"
)

func init() {
	kafka.InitKafka()
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalf("無法監聽 port: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, &handler.OrderServer{})

	fmt.Println("🚀 gRPC Server 已啟動於 port 8090")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("無法啟動 server: %v", err)
	}
}
