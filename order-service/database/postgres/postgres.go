package postgres

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	DBurl    string
}

func SetupDatabase(db *gorm.DB, models ...interface{}) error {
	err := db.AutoMigrate(models...)
	return err
}

func New(config *Config) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName,
	)

	if os.Getenv("ENV") == "staging" || os.Getenv("ENV") == "production" {
		db, err = gorm.Open(postgres.Open(config.DBurl), &gorm.Config{})
	} else {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
		return nil, err
	}

	return db, nil
}
