package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mhope-2/go-services/order-service/repository"
	"gorm.io/gorm"
)

type Handler struct {
	DB   *gorm.DB
	Repo *repository.Repository
}

func New(DB *gorm.DB) *Handler {
	repo := repository.New(DB)

	return &Handler{
		DB:   DB,
		Repo: repo,
	}
}

func (h *Handler) Register(v1 *gin.RouterGroup) {
	Orders := v1.Group("/orders")
	Orders.POST("/", h.CreateOrder)
	Orders.GET("/:id", h.RetrieveOrder)
}
