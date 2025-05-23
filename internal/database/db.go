package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	// в dsn вводим данные, которые мы указали при создании контейнера
	dsn := `host=localhost user=postgres password=123123 dbname=localdb port=5432 sslmode=disable`
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
}
