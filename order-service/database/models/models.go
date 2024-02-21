package models

import (
	"time"
)

type Order struct {
	ID               string    `gorm:"primary_key" json:"id"`
	UserID           string    `json:"user_id"`
	ProductCode      string    `json:"product_code"`
	CustomerFullName string    `json:"customer_full_name"`
	ProductName      string    `json:"product_name"`
	TotalAmount      float64   `json:"total_amount"`
	CreatedAt        time.Time `json:"created_at"`
}
