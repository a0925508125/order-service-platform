package repository

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type OrderRepository interface {
	AddStock(ctx context.Context, eventID, quantity int) error  //增加票數
	CheckStock(ctx context.Context, eventID int) (int64, error) //檢查庫存
	DecrStock(ctx context.Context, eventID, quantity int) error //扣除庫存

	// // 檢查用戶已購數量
	// CheckUserPurchased(ctx context.Context, eventID, userID int64) (int64, error)
	// // 累加用戶已購數量
	// IncrUserPurchased(ctx context.Context, eventID, userID int64, quantity int32) error
}

func NewOrderStockRepoRedis() OrderRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	log.Println("Redis initialized")

	return &OrderImp{rdb: rdb}
}

// func NewOrderMYSQLRepository() OrderRepository {
// 	return &mysqlOrderRepo{
// 		db: driver.GetDB(),
// 	}
// }
