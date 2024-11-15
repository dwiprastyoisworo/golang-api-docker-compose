package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Postgres struct {
	Host     string
	Username string
	Password string
	Database string
	Schema   string
	Port     string
	SslMode  string
}

type PostgresConnection struct {
	Postgres Postgres
}

func NewPostgresConnection(postgres Postgres) GetPostgresConnection {
	return &PostgresConnection{Postgres: postgres}
}

type GetPostgresConnection interface {
	Connect() *gorm.DB
}

func (p PostgresConnection) Connect() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		p.Postgres.Host, p.Postgres.Username, p.Postgres.Password, p.Postgres.Database, p.Postgres.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Postgres Connection Failed")
	}
	return db
}
