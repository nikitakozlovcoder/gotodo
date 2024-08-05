package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DatabaseConnectionString string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &Config{
		DatabaseConnectionString: os.Getenv("DATABASE_CONNECTION_STRING"),
	}, nil
}
