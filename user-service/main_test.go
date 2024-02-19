package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestGetUser tests the GetUser function for different user IDs
func TestGetUser(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Initialize router
	router := gin.Default()
	router.GET("/users/:id", GetUser)

	// Define test cases
	tests := []struct {
		description  string
		userID       string
		expectedCode int
		expectedBody string
	}{
		{"Valid user John Doe", "7c11e1ce2741", http.StatusOK, `{"id":"7c11e1ce2741","first_name":"John","last_name":"Doe"}`},
		{"Valid user Jane Doe", "e6f24d7d1c7e", http.StatusOK, `{"id":"e6f24d7d1c7e","first_name":"Jane","last_name":"Doe"}`},
		{"Unknown user", "unknown", http.StatusBadRequest, `{"message":"Unknown User"}`},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			// Create a request to pass to our handler
			req, _ := http.NewRequest("GET", "/users/"+test.userID, nil)

			// Create a response recorder to capture the response
			w := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(w, req)

			// Check the status code
			if w.Code != test.expectedCode {
				t.Errorf("Expected status code %d, got %d", test.expectedCode, w.Code)
			}

			// Check the response body
			if w.Body.String() != test.expectedBody {
				t.Errorf("Expected body %s, got %s", test.expectedBody, w.Body.String())
			}
		})
	}
}
