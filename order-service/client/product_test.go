package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestFetchProduct(t *testing.T) {
	// Set up a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example user data to return in the mock response
		mockProduct := Product{
			Code:  "product1",
			Name:  "Product 1",
			Price: 9.99,
		}

		resp, err := json.Marshal(mockProduct)
		if err != nil {
			t.Fatalf("Failed to marshal mock response: %v", err)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(resp)
		if err != nil {
			t.Fatalf("Failed to write mock response: %v", err)
		}
	}))
	defer server.Close()

	// mock user service url
	err := os.Setenv("PRODUCT_SERVICE_URL", server.URL)
	if err != nil {
		return
	}

	// Call FetchUserDetails
	product, err := FetchProduct("product1")
	if err != nil {
		t.Fatalf("FetchUser returned an error: %v", err)
	}

	// Verify the response
	if product.Code != "product1" || product.Name != "Product 1" || product.Price != 9.99 {
		t.Errorf("FetchProduct returned unexpected data: %+v", product)
	}
}
