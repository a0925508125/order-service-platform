package repository

import (
	"database/sql"
	"order-service-platform/worker/order/model"
)

type OrderMySQLimp struct {
	db *sql.DB
}

func (r *OrderMySQLimp) Create(order *model.OrderMessage) error {
	_, err := r.db.Exec(`
		INSERT INTO orders (order_id, event_id, user_id, quantity, status)
		VALUES (?, ?, ?, ?, ?)
	`,
		order.OrderID,
		order.EventID,
		order.UserID,
		order.Quantity,
		1,
	)
	return err
}
