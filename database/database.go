package database

import (
	"fmt"
	"hacktiv8-day12/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "localhost"
	PORT     = "5432"
	USER     = "postgres"
	PASS     = "Indra19"
	DB_NAME  = "final_assignment"
	APP_PORT = ":8888"
)

func ConnectDB() *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode= disable",
		HOST, PORT, USER, PASS, DB_NAME)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Default().Println("connected to database")

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Ping()

	// best practice, limit database connections or called connections pool
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxIdleTime(10 * time.Second)
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Photo{})
	db.AutoMigrate(models.Comment{})
	db.AutoMigrate(models.SocialMedia{})
	return db
}
