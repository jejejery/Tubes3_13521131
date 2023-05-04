package qnaController

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jejejery/src/backend/algorithm"
	database "github.com/jejejery/src/backend/db"
	"github.com/jejejery/src/backend/model"
)

func Index(c *fiber.Ctx) error {
	var qnas []model.QnA
	database.DB.Find(&qnas)

	return c.JSON(qnas)
}

func Create(c *fiber.Ctx) error {
	var request model.Request
	var qnas []model.QnA
	database.DB.Find(&qnas)
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request!",
		})
	}
	index := 0
	var similarityPercentage float32
	var idx int
	similarity := float32(0)
	for i := 0; i < len(qnas); i++ {
		if request.Algorithm {
			_, similarityPercentage = algorithm.KMPMatch(qnas[i].Question, request.Question)
		} else {
			_, similarityPercentage = algorithm.BMMatch(qnas[i].Question, request.Question)
		}
		if similarity < similarityPercentage {
			index = int(qnas[i].Id)
			similarity = similarityPercentage
			idx = i
		}
	}
	if similarity == 100 {

		//UPDATE
		var newUpdate struct {
			Id       uint
			Question string
			Answer   string
		}
		newUpdate.Id = uint(index)
		newUpdate.Question = request.Question
		newUpdate.Answer = request.Answer
		var temp float32
		if request.Algorithm {
			_, temp = algorithm.KMPMatch(newUpdate.Answer, qnas[idx].Answer)
		} else {
			_, temp = algorithm.BMMatch(newUpdate.Answer, qnas[idx].Answer)
		}
		if temp == 100 {
			return c.SendString("Pertanyaan " + request.Question + " dengan jawaban " + request.Answer + " sudah ada di database")
		}
		jsonStr, err := json.Marshal(newUpdate)
		if err != nil {
			return err
		}
		req, err := http.NewRequest("PUT", "http://localhost:8000/api/qna", bytes.NewBuffer((jsonStr)))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		defer resp.Body.Close()

		return c.SendString("Pertanyaan " + request.Question + " sudah ada di database dan jawaban diperbaharui menjadi " + request.Answer)
	} else {
		newQnA := model.QnA{Question: request.Question, Answer: request.Answer}
		database.DB.Create(&newQnA)
		return c.SendString("Pertanyaan " + request.Question + "dengan jawaban " + request.Answer + " berhasil ditambahkan ke database")
	}

}

func Delete(c *fiber.Ctx) error {
	var request model.Request
	var qnas []model.QnA
	database.DB.Find(&qnas)
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request!",
		})
	}
	index := 0
	var similarityPercentage float32
	var toDelete string
	similarity := float32(0)
	for i := 0; i < len(qnas); i++ {
		if request.Algorithm {
			_, similarityPercentage = algorithm.KMPMatch(qnas[i].Question, request.Question)
		} else {
			_, similarityPercentage = algorithm.BMMatch(qnas[i].Question, request.Question)
		}
		if similarity < similarityPercentage {
			index = int(qnas[i].Id)
			similarity = similarityPercentage
			toDelete = qnas[i].Question
		}
	}
	if similarity > 90 {
		var newDelete model.QnA
		database.DB.Where("Id = ?", index).Delete(&newDelete)
		return c.SendString("Pertanyaan " + toDelete + " berhasil dihapus!")
	} else {
		return c.SendString("Pertanyaan yang anda ingin dihapus tidak terdapat di database!")
	}

}

func Update(c *fiber.Ctx) error {
	var newUpdate struct {
		Id       uint
		Question string
		Answer   string
	}
	if err := c.BodyParser(&newUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request!",
		})
	}
	database.DB.Table("qn_as").Where("id= ?", newUpdate.Id).Updates(&newUpdate)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Data is successfully updated",
	})

}
