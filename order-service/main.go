package main

import (
	"fmt"
	"log"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mhope-2/go-services/order-service/database/models"
	"github.com/mhope-2/go-services/order-service/database/postgres"
	"github.com/mhope-2/go-services/order-service/handler"
	"github.com/mhope-2/go-services/order-service/server"
	"github.com/mhope-2/go-services/order-service/shared"
)

func main() {
	config := shared.NewEnvConfig()

	db, err := postgres.New(&postgres.Config{
		Host:     config.Host,
		Port:     config.Port,
		User:     config.User,
		Password: config.Password,
		DBName:   config.DBName,
		SSLMode:  config.SSLMode,
		DBurl:    config.DBurl,
	})

	if err != nil {
		log.Fatal("failed To Connect To Postgresql database", err)
	}

	err = postgres.SetupDatabase(db, &models.Order{})
	if err != nil {
		log.Fatalf("failed to setup tables, %v", err)
	}
	log.Println("DB migration completed")

	s := server.New()
	h := handler.New(db)

	routes := s.Group("")
	h.Register(routes)

	server.Start(&s, &server.Config{
		Port: fmt.Sprintf(":%s", config.ServerPort),
	})
}
