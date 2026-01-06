package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type OrderImp struct {
	rdb *redis.Client
}

func (s *OrderImp) AddStock(ctx context.Context, eventID, quantity int) error {
	key := fmt.Sprintf("stock:event:%d", eventID)
	return s.rdb.Set(ctx, key, quantity, 0).Err()
}

func (r *OrderImp) CheckStock(ctx context.Context, eventID int) (int64, error) {
	p := r.rdb.Ping(ctx)
	log.Println("ping ", p)
	return r.rdb.Get(ctx, fmt.Sprintf("stock:event:%d", eventID)).Int64()
}

func (r *OrderImp) DecrStock(ctx context.Context, eventID, quantity int) error {
	return r.rdb.DecrBy(ctx, fmt.Sprintf("stock:event:%d", eventID), int64(quantity)).Err()
}

// func (r *OrderImp) CheckUserPurchased(ctx context.Context, eventID, userID int64) (int64, error) {
// 	return r.rdb.Get(ctx, fmt.Sprintf("user:purchase:%d:%d", eventID, userID)).Int64()
// }

// func (r *OrderImp) IncrUserPurchased(ctx context.Context, eventID, userID int64, quantity int32) error {
// 	return r.rdb.IncrBy(ctx, fmt.Sprintf("user:purchase:%d:%d", eventID, userID), int64(quantity)).Err()
// }
