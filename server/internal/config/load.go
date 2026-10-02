package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseUrl string
}

func getPort() string {
	port := os.Getenv("HTTP_PORT")
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("HTTP_PORT is not an int: %v\n", err)
	}
	return port
}

func getDatabaseUrl() string {
	return ""
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error  loading .env file: %v\n", err)
	}
	return &Config{
		Port:        getPort(),
		DatabaseUrl: getDatabaseUrl(),
	}
}
