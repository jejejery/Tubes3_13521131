package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/jejejery/src/backend/db"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error to load .env files")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	app.Listen(":" + port)

}
