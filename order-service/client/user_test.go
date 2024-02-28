package client

import (
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

func TestFetchUser(t *testing.T) {
	// Set up a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example user data to return in the mock response
		mockUser := User{
			ID:        "7c11e1ce2741",
			FirstName: "John",
			LastName:  "Doe",
		}

		resp, err := json.Marshal(mockUser)
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
	err := os.Setenv("USER_SERVICE_URL", server.URL)
	if err != nil {
		return
	}

	// Call FetchUserDetails
	user, err := FetchUser("7c11e1ce2741")
	if err != nil {
		t.Fatalf("FetchUser returned an error: %v", err)
	}

	// Verify the response
	if user.ID != "7c11e1ce2741" || user.FirstName != "John" || user.LastName != "Doe" {
		t.Errorf("FetchUser returned unexpected data: %+v", user)
	}
}
