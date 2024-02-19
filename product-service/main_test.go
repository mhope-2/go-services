package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// SetupRouter initializes Gin router for testing
func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/products/:code", GetProduct)
	return router
}

// TestGetProduct tests the GetProduct function for different product codes
func TestGetProduct(t *testing.T) {
	router := SetupRouter()

	tests := []struct {
		description    string
		productCode    string
		expectedStatus int
		expectedBody   Product
	}{
		{
			description:    "Get product1",
			productCode:    "product1",
			expectedStatus: http.StatusOK,
			expectedBody:   Product{Code: "product1", Name: "Product 1", Price: 9.99},
		},
		{
			description:    "Get product2",
			productCode:    "product2",
			expectedStatus: http.StatusOK,
			expectedBody:   Product{Code: "product2", Name: "Product 2", Price: 14.99},
		},
		{
			description:    "Get non-existing product",
			productCode:    "non-existing",
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   Product{},
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/products/"+test.productCode, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, test.expectedStatus, w.Code)

			if w.Code == http.StatusOK {
				var product Product
				err := json.Unmarshal(w.Body.Bytes(), &product)
				assert.Nil(t, err)
				assert.Equal(t, test.expectedBody, product)
			} else {
				assert.Contains(t, w.Body.String(), "Internal Server Error")
			}
		})
	}
}
