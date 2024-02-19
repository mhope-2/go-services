package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type Product struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Define a GET route
	router.GET("/products/:code", GetProduct)

	err := router.Run(":8082")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func GetProduct(c *gin.Context) {
	var code = c.Param("code")

	switch code {

	case "product1":
		time.Sleep(300 * time.Millisecond) // Simulate some processing delay

		c.JSON(http.StatusOK, Product{
			Code:  "product1",
			Name:  "Product 1",
			Price: 9.99,
		})

	case "product2":
		time.Sleep(60 * time.Second)

		c.JSON(http.StatusOK, Product{
			Code:  "product2",
			Name:  "Product 2",
			Price: 14.99,
		})

	case "product3":
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
}
