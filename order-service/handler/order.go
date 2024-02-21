package handler

import (
	"log"
	"net/http"

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

	if err := validate.Struct(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Validation failed",
		})
		return
	}

	order, err := h.Repo.CreateOrder(data)
	if err != nil {
		log.Println("Error creating order: ", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"order": order,
	})
}

func (h *Handler) RetrieveOrder(c *gin.Context) {
	id := c.Param("id")

	order, err := h.Repo.RetrieveOrder(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed To retrieve customer",
		})
		return
	}

	if order == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"detail": "Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"order": order,
	})
	return
}
