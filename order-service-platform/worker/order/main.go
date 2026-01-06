package main

import (
	"context"
	"order-service-platform/kafka"
	"order-service-platform/worker/order/handler"
)
// netsh advfirewall set allprofiles state off

func main() {

	h := handler.NewOrderHandler()

	kafka.NewConsumer(
		kafka.Brokers,
		kafka.TopicOrder,
		kafka.GroupOrder,
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	kafka.ConsumeMessages(ctx, h.Handle)
	h.Create(ctx)
	select {}
	// v	shutdown.Wait()
}
