package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var JwtSecret string
var PostgresUri string
var MongoUri string
var Gin_mode string

func Load() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgres_user := os.Getenv("POSTGRES_USER")
	if postgres_user == "" {
		log.Fatal("Error on get POSTGRES_USER")
	}
	postgres_password := os.Getenv("POSTGRES_PASSWORD")
	if postgres_password == "" {
		log.Fatal("Error on get POSTGRES_PASSWORD")
	}
	postgres_db := os.Getenv("POSTGRES_DB")
	if postgres_db == "" {
		log.Fatal("Error on get POSTGRES_DB")
	}
	postgres_port := os.Getenv("POSTGRES_PORT")
	if postgres_port == "" {
		log.Fatal("Error on get POSTGRES_PORT")
	}
	postgres_host := os.Getenv("POSTGRES_HOST")
	if postgres_host == "" {
		log.Fatal("Error on get POSTGRES_HOST")
	}
	postgres_ssl := os.Getenv("POSTGRES_SSL")
	if postgres_ssl == "" {
		log.Fatal("Error on get POSTGRES_SSL")
	}

	PostgresUri = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", postgres_user, postgres_password, postgres_host, postgres_port, postgres_db, postgres_ssl)

	JwtSecret = os.Getenv("JWT_SECRET")
	if JwtSecret == "" {
		log.Fatal("Error on get JWT_SECRET")
	}

	Gin_mode = os.Getenv("GIN_MODE")
	if Gin_mode == "" {
		log.Fatal("Error on get GIN_MODE")
	}
}
