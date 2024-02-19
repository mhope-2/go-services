package repository

import (
	"github.com/google/uuid"
	"github.com/mhope-2/go-services/order-service/shared"
	"time"

	"github.com/mhope-2/go-services/order-service/database/models"
)

type Order interface {
	ListOrders() ([]models.Order, error)
	CreateOrder(data shared.CreateOrderRequest) (*models.Order, error)
	RetrieveOrder(id string) (*models.Order, error)
}

func (r *Repository) ListOrders() ([]models.Order, error) {
	var orders []models.Order

	if err := r.DB.Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *Repository) CreateOrder(data shared.CreateOrderRequest) (*models.Order, error) {
	order := models.Order{
		ID:               uuid.New().String(),
		UserID:           data.UserID,
		ProductCode:      data.ProductCode,
		CustomerFullName: data.CustomerFullName,
		ProductName:      data.ProductName,
		TotalAmount:      data.TotalAmount,
		CreatedAt:        time.Now(),
	}

	if err := r.DB.Create(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil

}

func (r *Repository) RetrieveOrder(id string) (*models.Order, error) {

	order := models.Order{}

	if err := r.DB.First(&order, id).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
