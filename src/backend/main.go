package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jejejery/src/backend/controller/historyController"
	inputcontroller "github.com/jejejery/src/backend/controller/inputController"
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
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))

	api := app.Group("/api")
	qna := api.Group("/qna")
	input := api.Group("/input")
	history := api.Group("/history")

	qna.Get("/", qnaController.Index)
	qna.Get("/", qnaController.Show)
	qna.Post("/", qnaController.Create)
	// qna.Put("/:id", qnaController.Update)

	qna.Delete("/", qnaController.Delete)

	history.Get("/", historyController.Index)
	history.Get("/:id", historyController.Show)
	history.Post("/", historyController.Create)
	// qna.Put("/:id", qnaController.Update)

	history.Delete("/:id", historyController.Delete)

	input.Get("/", inputcontroller.Index)
	input.Get("/:id", inputcontroller.Show)
	input.Post("/", inputcontroller.Create)
	// qna.Put("/:id", qnaController.Update)

	input.Delete("/:id", inputcontroller.Delete)

	app.Listen(":" + port)

}
