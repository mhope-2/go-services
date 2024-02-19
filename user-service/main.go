package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	"time"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var (
	toggleStatusCode bool // Starts as false and will be toggled
	mu               sync.Mutex
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Define a GET route
	router.GET("/users/:id", GetUser)

	err := router.Run(":8081")
	if err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

// GetUser returns user object
func GetUser(c *gin.Context) {
	var ID = c.Param("id")

	switch ID {

	case "7c11e1ce2741":
		time.Sleep(300 * time.Millisecond) // Simulate some processing delay

		c.JSON(http.StatusOK, User{
			ID:        "7c11e1ce2741",
			FirstName: "John",
			LastName:  "Doe",
		})

	case "e6f24d7d1c7e":
		time.Sleep(300 * time.Millisecond)

		mu.Lock() // Ensure thread-safe access to toggleStatusCode
		if toggleStatusCode {
			c.JSON(http.StatusOK, User{
				ID:        "e6f24d7d1c7e",
				FirstName: "Jane",
				LastName:  "Doe",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}

		toggleStatusCode = !toggleStatusCode // Toggle the status code for next call
		mu.Unlock()

	default:
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unknown User"})
	}
}
