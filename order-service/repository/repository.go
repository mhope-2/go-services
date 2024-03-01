// Package repository maintains the data access components
package repository

import (
	"github.com/mhope-2/go-services/order-service/client"
	"gorm.io/gorm"
)

type Repository struct {
	DB             *gorm.DB
	UserService    client.UserService
	ProductService client.ProductService
}

func New(db *gorm.DB, userService client.UserService, productService client.ProductService) *Repository {
	return &Repository{db, userService, productService}
}
