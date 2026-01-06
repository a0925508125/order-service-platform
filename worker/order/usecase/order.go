package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"order-service-platform/worker/order/model"
	"order-service-platform/worker/order/repository"

	"github.com/go-redis/redis/v8"
)

type OrderUsecase struct {
	stockRepo repository.OrderRepository
	orderRepo repository.OrderRepoInterface
}

func NewOrderUsecase() *OrderUsecase {
	stockRepo := repository.NewOrderStockRepoRedis()
	orderRepo := repository.NewOrderRepoMySQL()
	return &OrderUsecase{
		stockRepo: stockRepo,
		orderRepo: orderRepo,
	}
}
func (u *OrderUsecase) AddStock(ctx context.Context, eventID, quantity int) error {
	return u.stockRepo.AddStock(ctx, eventID, quantity)
}

// 高併發下單邏輯
func (u *OrderUsecase) ProcessOrder(ctx context.Context, order *model.OrderMessage) error {
	remain, err := u.stockRepo.CheckStock(ctx, order.EventID)
	if err != nil {
		if err == redis.Nil {
			// key 不存在
			fmt.Println("stock not set yet")
		}
		log.Println("ProcessOrder CheckStock", err)
		return err
	}

	if remain < int64(order.Quantity) {
		// 庫存不足
		return errors.New("remain 0")
	}

	// 扣除數量
	err = u.stockRepo.DecrStock(ctx, order.EventID, order.Quantity)
	if err != nil {
		log.Println("ProcessOrder DecrStock", err)
		return err
	}

	//寫入訂單
	err = u.orderRepo.Create(order)
	if err != nil {
		log.Println("ProcessOrder DecrStock", err)
		return err
	}

	return nil
}
