package database

import (
	"log"
	"os"

	"github.com/jejejery/src/backend/model"
	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load .env file!")
	}
	dsn := os.Getenv("DSN")
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database")
	} else {
		log.Println("Connect succesfully!")
	}
	DB = database

	database.AutoMigrate(
		&model.QnA{},
		&model.History{},
	)

}
