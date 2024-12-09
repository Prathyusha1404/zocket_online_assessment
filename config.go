package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds application configuration for the database and Redis
type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
	Redis struct {
		Address  string
		Password string
	}
}

// LoadConfig loads configuration from the .env file
func LoadConfig() *Config {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Convert database port to an integer
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Fatalf("Invalid port value in .env file: %v", err)
	}

	// Return the configuration struct
	return &Config{
		Database: struct {
			Host     string
			Port     int
			User     string
			Password string
			Name     string
		}{
			Host:     os.Getenv("DATABASE_HOST"),
			Port:     port,
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Name:     os.Getenv("DATABASE_NAME"),
		},
		Redis: struct {
			Address  string
			Password string
		}{
			Address:  os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
	}
}
