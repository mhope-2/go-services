package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/mhope-2/go-services/order-service/shared"
)

// User holds the structure of the user response object
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UserService interface {
	FetchUser(userID string) (*User, error)
}

type HTTPUserService struct{}

// FetchUser asynchronously makes a request to the user service and returns the user object with retry mechanism
func (s *HTTPUserService) FetchUser(userID string) (*User, error) {
	config := shared.NewEnvConfig()

	url := fmt.Sprintf("%s/%s", config.UserServiceURL, userID)

	// Channel to receive the result
	resultChan := make(chan *User)
	errorChan := make(chan error)

	go func() {
		var resp *http.Response
		var err error

		// make request with retries to the user service
		retries := 3
		for i := 0; i < retries; i++ {
			resp, err = http.Get(url)
			if err != nil {
				errorChan <- fmt.Errorf("request to user service failed: %w", err)
				return
			}
			if resp.StatusCode == http.StatusOK {
				break
			}
			log.Printf("Retry %d/3: received non-200 status code from user service: %d\n\n", i+1, resp.StatusCode)
			time.Sleep(300 * time.Millisecond) // Wait before retrying
		}

		defer func() {
			err = resp.Body.Close()
			if err != nil {
				log.Println("failed to close user response body")
			}
		}()

		// Check for non-200 (OK) status code
		if resp.StatusCode != http.StatusOK {
			errorChan <- fmt.Errorf("received non-200 status code from user service: %d", resp.StatusCode)
			return
		}

		// Read and parse the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errorChan <- fmt.Errorf("failed to read response body: %w", err)
			return
		}

		var user User
		if err = json.Unmarshal(body, &user); err != nil {
			errorChan <- fmt.Errorf("failed to unmarshal response body: %w", err)
			return
		}

		resultChan <- &user
	}()

	select {
	case user := <-resultChan:
		return user, nil
	case err := <-errorChan:
		return nil, err
	case <-time.After(5 * time.Second):
		return nil, fmt.Errorf("request to user service timed out")
	}
}
