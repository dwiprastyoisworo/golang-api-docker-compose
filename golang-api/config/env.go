package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Env struct {
	Postgres Postgres
	Port     string
}

var GlobalEnv Env

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgres := Postgres{
		Host:     os.Getenv("POSTGRES_DB_HOST"),
		Username: os.Getenv("POSTGRES_DB_USERNAME"),
		Password: os.Getenv("POSTGRES_DB_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB_NAME"),
		Schema:   os.Getenv("POSTGRES_DB_SCHEMA"),
		Port:     os.Getenv("POSTGRES_DB_PORT"),
		SslMode:  os.Getenv("POSTGRES_DB_SSL"),
	}

	GlobalEnv.Postgres = postgres
	GlobalEnv.Port = os.Getenv("PORT")
}
