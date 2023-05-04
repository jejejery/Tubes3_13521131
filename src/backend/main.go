package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	inputcontroller "github.com/jejejery/src/backend/controller/inputController"
	"github.com/jejejery/src/backend/controller/qnaController"
	"github.com/jejejery/src/backend/controller/sessionController"
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
	session := api.Group("/session")

	qna.Get("/", qnaController.Index)
	qna.Post("/", qnaController.Create)
	qna.Put("/", qnaController.Update)
	qna.Delete("/", qnaController.Delete)

	input.Get("/", inputcontroller.Index)
	input.Get("/toShow", inputcontroller.Show)
	input.Post("/", inputcontroller.Create)

	session.Get("/", sessionController.Index)
	session.Post("/", sessionController.Create)

	app.Listen(":" + port)

}
