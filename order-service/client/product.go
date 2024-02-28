package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/mhope-2/go-services/order-service/shared"
)

// Product holds the structure of the product response object
type Product struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// FetchProduct asynchronously makes a request to the product details service and returns the product details
func FetchProduct(code string) (*Product, error) {
	config := shared.NewEnvConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resultChan := make(chan *Product, 1)
	errorChan := make(chan error, 1)

	go func() {
		url := fmt.Sprintf("%s/%s", config.ProductServiceURL, code)
		req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
		if err != nil {
			errorChan <- fmt.Errorf("failed to create product request: %w", err)
			return
		}

		// Make the request
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			errorChan <- fmt.Errorf("request to product service failed: %w", err)
			return
		}
		defer func() {
			err := resp.Body.Close()
			if err != nil {
				log.Println("failed to close product response body")
			}
		}()

		if resp.StatusCode == http.StatusInternalServerError {
			errorChan <- fmt.Errorf("product service returned a server error: %d", resp.StatusCode)
			return
		} else if resp.StatusCode != http.StatusOK {
			errorChan <- fmt.Errorf("received non-200 status code from product service: %d", resp.StatusCode)
			return
		}

		// Read and parse the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errorChan <- fmt.Errorf("failed to read response body: %w", err)
			return
		}

		var product Product
		if err := json.Unmarshal(body, &product); err != nil {
			errorChan <- fmt.Errorf("failed to unmarshal response body: %w", err)
			return
		}

		resultChan <- &product
	}()

	select {
	case product := <-resultChan:
		return product, nil
	case err := <-errorChan:
		return nil, err
	case <-ctx.Done():
		return nil, fmt.Errorf("request to product service timed out")
	}
}
