package qnaController

import (
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

func Create(c *fiber.Ctx) error {
	var newQna model.QnA
	var qna model.QnA
	if err := c.BodyParser(&newQna); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request!!",
		})
	}
	if err := database.DB.Where("question = ?", newQna.Question).First(&qna).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := database.DB.Create(&newQna).Error; err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal server error!!",
				})
			}
			return c.JSON(newQna)
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error!!",
			})
		}
	}
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": "The question is available",
	})

}

func Delete(c *fiber.Ctx) error {
	var deleteQna model.QnA
	var temp model.QnA
	if err := c.BodyParser(&deleteQna); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request!!",
		})
	}
	if err := database.DB.Where("question = ?", deleteQna.Question).First(&temp).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "The question you want to delete is unavailable! You can't delete it",
		})
	} else {
		database.DB.Where("question = ?", deleteQna.Question).Delete(&temp)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Record deleted successfully",
		})
	}
}
