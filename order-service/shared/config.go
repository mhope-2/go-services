// Package shared maintains the group of shared types and configs
package shared

import (
	"os"
)

// EnvConfig represents the application configuration - Env variables
type EnvConfig struct {
	User              string
	Password          string
	DBName            string
	SSLMode           string
	Host              string
	Port              string
	DBurl             string
	ServerPort        string
	AMQPURI           string
	UserServiceURL    string
	ProductServiceURL string
}

// NewEnvConfig creates a new Config instance, loading values from environment variables
func NewEnvConfig() *EnvConfig {
	return &EnvConfig{
		User:              os.Getenv("DB_USER"),
		Password:          os.Getenv("DB_PASS"),
		DBName:            os.Getenv("DB_NAME"),
		SSLMode:           os.Getenv("DB_SSLMODE"),
		Host:              os.Getenv("DB_HOST"),
		Port:              os.Getenv("DB_PORT"),
		DBurl:             os.Getenv("DB_URL"),
		ServerPort:        os.Getenv("PORT"),
		AMQPURI:           os.Getenv("AMQP_URI"),
		UserServiceURL:    os.Getenv("USER_SERVICE_URL"),
		ProductServiceURL: os.Getenv("PRODUCT_SERVICE_URL"),
	}
}
