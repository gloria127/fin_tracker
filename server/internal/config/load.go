package config

import (
	"fmt"
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
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	//обязательные параметры
	if host == "" {
		log.Fatal("POSTGRES_HOST is empty")
	}

	if user == "" {
		log.Fatal("POSTGRES_USER is empty")
	}

	if password == "" {
		log.Fatal("POSTGRES_PASSWORD is empty")
	}

	if dbname == "" {
		log.Fatal("POSTGRES_DB is empty")
	}

	portNum, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalf("POSTGRES_PORT is not an int: %v", err)
	}

	if portNum < 1 || portNum > 65535 {
		log.Fatal("POSTGRES_PORT must be between 1 and 65535")
	}

	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, portNum, user, password, dbname)
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
