package repository

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mhope-2/go-services/order-service/database/postgres"
	"github.com/mhope-2/go-services/order-service/messaging/rabbitmq"
	"github.com/mhope-2/go-services/order-service/shared"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Setenv("AMQP_URI", "amqp://user:order@localhost:5672/")

	os.Exit(m.Run())
}

func TestOrderRepository(t *testing.T) {
	t.Run("crud", crud)
}

func crud(t *testing.T) {

	db, err := postgres.New(&postgres.Config{
		Host:     "localhost",
		Port:     "5433",
		User:     "postgres",
		Password: "postgres",
		DBName:   "orders",
		SSLMode:  "disable",
	})

	if err != nil {
		log.Fatal("failed To Connect To Postgresql database", db, err)
	}

	repo := New(db)
	_, err = rabbitmq.NewPublisher("amqp://user:order@localhost:5672/")
	if err != nil {
		fmt.Println("failed to create amqp connection")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	data := shared.CreateOrderRequest{UserID: "7c11e1ce2741", ProductCode: "product1"}
	order, err := repo.CreateOrder(ctx, data)

	if err != nil {
		t.Fatalf("Error creating url: %v", err)
	}

	assert.NoError(t, err, "Failed to create order")
	assert.Equal(t, order.UserID, data.UserID)
	assert.Equal(t, order.ProductCode, data.ProductCode)

	// ------------------------------------------------------------------------

	//// test repo url retrieve
	//url, err = repo.GetUrl("A1B2C3D4", "127.0.0.1")
	//if err != nil {
	//	t.Fatalf("Error retrieving url: %v", err)
	//}
	//
	//assert.NoError(t, err, "Failed to retrieve url")
	//assert.Equal(t, url.Url, "https://youtube.com/")
	//assert.Equal(t, url.Slug, "A1B2C3D4")
}
