package handler

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mhope-2/go-services/order-service/shared"
)

var validate *validator.Validate

func (h *Handler) CreateOrder(c *gin.Context) {
	var data shared.CreateOrderRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"detail": "failed to parse request",
		})
		return
	}

	validate = validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation failed",
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	order, err := h.Repo.CreateOrder(ctx, data)
	if err != nil {
		log.Println("Error creating order: ", err)
		c.JSON(http.StatusOK, gin.H{
			"detail": "order creation failed",
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *Handler) RetrieveOrder(c *gin.Context) {
	id := c.Param("id")

	order, err := h.Repo.RetrieveOrder(id)
	if order == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"detail": "Not Found",
		})
		return
	}

	if err != nil {
		log.Println("failed to retrieve order: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed To retrieve customer",
		})
		return
	}

	c.JSON(http.StatusOK, order)
	return
}
