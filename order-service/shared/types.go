package shared

import "time"

type CreateOrderRequest struct {
	UserID      string `json:"user_id" validate:"required"`
	ProductCode string `json:"product_code" validate:"required"`
}

type Order struct {
	OrderID          string    `json:"order_id"`
	CustomerFullName string    `json:"customer_fullname"`
	ProductName      string    `json:"product_name"`
	TotalAmount      float64   `json:"total_amount"`
	CreatedAt        time.Time `json:"created_at"`
}

type Message struct {
	Producer string           `json:"producer"`
	SentAt   time.Time        `json:"sent_at"`
	Type     string           `json:"type"`
	Payload  map[string]Order `json:"payload"`
}
