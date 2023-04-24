package inputcontroller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	database "github.com/jejejery/src/backend/db"
	"github.com/jejejery/src/backend/model"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var input []model.InputUser
	database.DB.Find(&input)
	return c.JSON(input)
}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")
	var input model.InputUser
	if err := database.DB.First(&input, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data is not found!",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data is not found!",
		})
	}
	database.DB.First(&input, id) // ambil data pertamma yang match
	return c.JSON(input)
}

func Create(c *fiber.Ctx) error {
	var newInput model.InputUser
	if err := c.BodyParser(&newInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request!!",
		})
	}
	if err := database.DB.Create(&newInput).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error!!",
		})
	}
	return c.JSON(newInput)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	var input model.InputUser
	if database.DB.Delete(&input, id).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Can not delete the data",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data is successfully deleted!",
	})
}
