package order

import (
	"context"
)

var ctx = context.Background()

type OrderUsecase struct {
}

func NewTicketService() *OrderUsecase {
	return &OrderUsecase{}
}
