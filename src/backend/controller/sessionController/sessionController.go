package sessionController

import (
	"github.com/gofiber/fiber/v2"
	database "github.com/jejejery/src/backend/db"
	"github.com/jejejery/src/backend/model"
)

func Index(c *fiber.Ctx) error {
	var sessions []model.Sessions
	database.DB.Find(&sessions)

	return c.JSON(sessions)
}

func Create(c *fiber.Ctx) error {

	var input struct {
		Session int64
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request!!",
		})
	}
	newInput := model.Sessions{Session: input.Session}
	if err := database.DB.Create(&newInput).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error!!",
		})
	}
	return c.JSON(input)

}
