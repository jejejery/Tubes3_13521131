package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jejejery/src/backend/controller/qnaController"
	database "github.com/jejejery/src/backend/db"
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

	api := app.Group("/api")
	qna := api.Group("/qna")
	// history := api.Group("/history")
	qna.Get("/", qnaController.Index)
	qna.Get("/:id", qnaController.Show)
	qna.Post("/", qnaController.Create)
	// qna.Put("/:id", qnaController.Update)
	qna.Delete("/:id", qnaController.Delete)

	// history.Get("/", historyController.Index)
	// history.Get("/:id", historyController.Show)
	// history.Post("/", historyController.Create)
	// history.Put("/:id", historyController.Update)
	// history.Delete("/:id", historyController.Delete)

	app.Listen(":" + port)

}
