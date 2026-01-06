package model

type OrderMessage struct {
	OrderID   string `json:"orderID"` //uuid
	UserID    int64  `json:"userID"`
	EventID   int    `json:"eventID"`  //下標品項
	Quantity  int    `json:"quantity"` //下單數量
	Timestamp int64  `json:"timestamp"`
}
