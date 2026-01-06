package model

type Order struct {
	EventID  int64 `json:"eventId"`
	UserID   int64 `json:"userId"`
	Quantity int32 `json:"quantity"`
}
