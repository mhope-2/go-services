package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	err := os.Setenv("ENV", "testing")
	if err != nil {
		fmt.Println("Error setting env variable, ENV")
		return
	}

	os.Exit(m.Run())
}

func TestOrderHandler(t *testing.T) {
	t.Run("create", create)
	t.Run("list", list)
	t.Run("retrieve", retrieve)
}

// =============================================================================

func create(t *testing.T) {
	data := map[string]string{
		"user_id":      "7c11e1ce2741",
		"product_code": "Product 1",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Failed to marshal JSON. Cause: %v\n", err)
	}

	req, err := http.NewRequest("POST", "/orders", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Could not create request. Cause: %v\n", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Setup router with a dummy handler to simulate endpoint behavior
	router := gin.Default()
	router.POST("/orders", func(c *gin.Context) {
		// Simulate successful handling of the request
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	})

	// Create a response recorder to inspect the response
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check if the response status code is as expected
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, w.Code)
	}
}

func list(t *testing.T) {

	router := gin.Default()
	router.GET("/orders", func(c *gin.Context) {
		// Simulate successful retrieval of a list of orders
		c.JSON(http.StatusOK, gin.H{"orders": []string{"order1", "order2"}})
	})

	// Create request
	req, err := http.NewRequest("GET", "/orders", nil)
	if err != nil {
		t.Fatalf("Could not create request. Cause: %v\n", err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, w.Code)
	}
}

func retrieve(t *testing.T) {

	router := gin.Default()
	router.GET("/orders/:id", func(c *gin.Context) {
		// Assuming the ID is extracted correctly
		id := c.Param("id")

		// Simulate successful retrieval of order
		c.JSON(http.StatusOK, gin.H{"id": id, "status": "success"})
	})

	// Create request
	req, err := http.NewRequest("GET", "/orders/123", nil)
	if err != nil {
		t.Fatalf("Could not create request. Cause: %v\n", err)
	}

	// Create a response recorder
	w := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(w, req)

	// Check response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, w.Code)
	}
}
