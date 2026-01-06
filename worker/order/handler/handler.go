package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"order-service-platform/worker/order/model"
	"order-service-platform/worker/order/usecase"
)

type OrderHandler struct {
	uc *usecase.OrderUsecase
}

func NewOrderHandler() *OrderHandler {
	return &OrderHandler{
		uc: usecase.NewOrderUsecase(),
	}
}

type OrderMessage struct {
	OrderId   string `json:"orderId"`
	EventId   int64  `json:"eventId"`
	UserId    int64  `json:"userId"`
	Quantity  int32  `json:"quantity"`
	Timestamp int64  `json:"timestamp"`
}

func (h *OrderHandler) Create(ctx context.Context) error {
	err := h.uc.AddStock(ctx, 2, 100)
	if err != nil {
		log.Println("err:", err)
	}
	return nil
}

func (h *OrderHandler) Handle(ctx context.Context, msg []byte) error {
	var order model.OrderMessage
	if err := json.Unmarshal(msg, &order); err != nil {
		log.Println("json unmarshal error:", err)
		return err
	}

	fmt.Printf("[Handler] Received order ID:%s, msg:%+v", order.OrderID, order)

	return h.uc.ProcessOrder(ctx, &order)
}
