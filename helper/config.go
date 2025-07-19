package helper

import (
	"log"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	// Path to the config file
	MONGO_DB_CONNECTION_STRING string
}

func NewConfig() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading .env file")
	}
	return &Config{
		MONGO_DB_CONNECTION_STRING: "mongodb://localhost:27017",
	}
}
