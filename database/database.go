package database

import (
	"log"
	"time"
	"web-api-go/database/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	str := "host=localhost port=5432 user=admin dbname=books sslmode=disable  password=123456 "

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	db = database
	migrations.RunMigrations(database)
	config, _ := db.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDB() *gorm.DB {
	return db
}
