package qnaController

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	database "github.com/jejejery/src/backend/db"
	"github.com/jejejery/src/backend/model"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var qnas []model.QnA
	database.DB.Find(&qnas)
	return c.JSON(qnas)
}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")
	var qna model.QnA
	if err := database.DB.First(&qna, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data is not found!",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Data is not found!",
		})
	}
	database.DB.First(&qna, id) // ambil data pertamma yang match
	return c.JSON(qna)
}

func Create(c *fiber.Ctx) error {
	var newQna model.QnA
	if err := c.BodyParser(&newQna); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request!!",
		})
	}
	if err := database.DB.Create(&newQna).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error!!",
		})
	}
	return c.JSON(newQna)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	var qna model.QnA
	if database.DB.Delete(&qna, id).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Can not delete the data",
		})
	}
	return c.JSON(fiber.Map{
		"message" : "Data is successfully deleted!",
	})
}
