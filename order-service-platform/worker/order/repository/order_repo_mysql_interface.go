package repository

import (
	"database/sql"
	"log"
	"order-service-platform/worker/order/model"

	_ "github.com/go-sql-driver/mysql"
)

type OrderRepoInterface interface {
	Create(order *model.OrderMessage) error
}

func NewOrderRepoMySQL() OrderRepoInterface {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/shop")
	if err != nil {
		log.Fatal("MySQL init error:", err)
	}
	log.Println("MySQL initialized")
	return &OrderMySQLimp{db: db}
}
