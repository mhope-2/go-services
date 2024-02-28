package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/mhope-2/go-services/order-service/client"
	"github.com/mhope-2/go-services/order-service/messaging/rabbitmq"
	"github.com/mhope-2/go-services/order-service/shared"
	"log"
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

func (r *Repository) CreateOrder(ctx context.Context, data shared.CreateOrderRequest) (*models.Order, error) {

	userChan := make(chan *client.User, 1)
	productChan := make(chan *client.Product, 1)
	errorChan := make(chan error, 2)

	go func() {
		user, err := client.FetchUser(data.UserID)
		if err != nil {
			errorChan <- fmt.Errorf("failed to fetch user details: %w", err)
			return
		}
		userChan <- user
	}()

	go func() {
		product, err := client.FetchProduct(data.ProductCode)
		if err != nil {
			errorChan <- fmt.Errorf("failed to fetch product details: %w", err)
			return
		}
		productChan <- product
	}()

	var user *client.User
	var product *client.Product

	for i := 0; i < 2; i++ {
		select {
		case user = <-userChan:
		case product = <-productChan:
		case err := <-errorChan:
			log.Printf("%s", err)
			return nil, err
		}
	}

	order := models.Order{
		ID:               uuid.New().String(),
		UserID:           data.UserID,
		ProductCode:      data.ProductCode,
		CustomerFullName: fmt.Sprintf("%s %s", user.FirstName, user.LastName),
		ProductName:      product.Name,
		TotalAmount:      product.Price,
		CreatedAt:        time.Now(),
	}

	// Use a transaction in order creation
	tx := r.DB.Begin()

	if err := tx.WithContext(ctx).Create(&order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("commit create order transaction failed: %w", err)
	}

	// create message and publish to RabbitMQ
	orderPayload := shared.Order{
		OrderID:          order.ID,
		CustomerFullName: order.CustomerFullName,
		ProductName:      order.ProductName,
		TotalAmount:      order.TotalAmount,
		CreatedAt:        order.CreatedAt,
	}

	message := shared.Message{
		Producer: "order",
		SentAt:   time.Now(),
		Type:     "created_order",
		Payload: map[string]shared.Order{
			"order": orderPayload,
		},
	}

	rabbitmq.Publish(message, "order_queue", "created_order", "orders")

	return &order, nil

}

func (r *Repository) RetrieveOrder(id string) (*models.Order, error) {

	order := models.Order{}

	if err := r.DB.Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}

	return &order, nil
}
