package grpcclient

import (
	"order-service-platform/proto/proto/pb"

	"google.golang.org/grpc"
)

var (
	OrderClient pb.OrderServiceClient
)

func Init() {
	conn, err := grpc.Dial(
		"127.0.0.1:8090",
		grpc.WithInsecure())

	if err != nil {
		panic(err.Error())
	}

	OrderClient = pb.NewOrderServiceClient(conn)
}
