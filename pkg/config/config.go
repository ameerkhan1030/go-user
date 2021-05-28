package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config contains the environment specific configuration values needed by the
// application.
type Config struct {
	Environment               string
	DatabaseHost              string
	DatabaseName              string
	DatabasePassword          string
	DatabasePort              int
	DatabaseUser              string
	BasicAuthUsername string
	BasicAuthPassword  string

	Port					  int
}

// New returns an instance of Config based on the "ENVIRONMENT" environment
// variable.
func New() (Config, error) {
	err := godotenv.Load(".env")

  if err != nil {
    log.Fatalf("Error loading .env file")
  }
	cfg := Config{
		DatabasePort:              3306,
		DatabaseHost: os.Getenv("DATABASE_HOST"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseUser: os.Getenv("DATABASE_USER"),
		BasicAuthUsername: os.Getenv("AUTH_USERNAME"),
		BasicAuthPassword: os.Getenv("AUTH_PASSWORD"),
		Port: 3001,
	}
	return cfg, nil
}
