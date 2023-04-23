package historyController

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	database "github.com/jejejery/src/backend/db"
	"github.com/jejejery/src/backend/model"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var histories []model.History
	database.DB.Find(&histories)
	return c.JSON(histories)
}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")
	var history model.History
	if err := database.DB.First(&history, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data is not found!",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data is not found!",
		})
	}
	database.DB.First(&history, id) // ambil data pertamma yang match
	return c.JSON(history)
}

func Create(c *fiber.Ctx) error {
	var newHistory model.History
	if err := c.BodyParser(&newHistory); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request!!",
		})
	}
	if err := database.DB.Create(&newHistory).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error!!",
		})
	}
	return c.JSON(newHistory)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	var history model.History
	if database.DB.Delete(&history, id).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Can not delete the data",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data is successfully deleted!",
	})
}
