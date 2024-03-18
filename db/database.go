package db

import (
	"emp-mini/entity"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDB() {
	time.Sleep(10 * time.Second)
	// db, err = gorm.Open(postgres.Open("postgresql://postgres:prithvi@db-service:5432/minidb?sslmode=disable"), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("failed to connect database")
	// }

	dsn := "host=postgres user=postgres password=prithvi dbname=minidb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// AutoMigrate the Employee schema
	err = db.AutoMigrate(&entity.Employee{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}

	// AutoMigrate the Rating schema
	err = db.AutoMigrate(&entity.Rating{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}
}
func GetDB() *gorm.DB {
	return db
}
