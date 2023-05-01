package inputcontroller

import (
	"net/http"
	"sort"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jejejery/src/backend/algorithm"
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
	var input struct {
		Input     string `json:"Input"`
		Algorithm bool   `json:"Algorithm"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request!!",
		})
	}
	answer := algorithm.CheckQuestion(input.Input)
	if answer == "" {
		var qnas []model.QnA
		database.DB.Find(&qnas)
		isMatch := false
		// save similarity match for match n non match pattern so far
		matchAnswerString := ""
		similarityMatch := 0.0
		similarityNonMatch := 0.0
		nonMatchStringQuestion := ""
		nonMatchStringAnswer := ""
		type stringSimilarity struct {
			nonMatchStringQuestion string
			nonMatchStringAnswer   string
			similarityPercentage   float64
		}
		var nonMatchStrings []stringSimilarity
		n := len(qnas)
		for i := 0; i < n; i++ {
			if input.Algorithm {
				// kmp
				index, similarityTemp := (algorithm.KMPMatch(input.Input, qnas[i].Question))
				println(similarityTemp)
				if index == -1 && similarityTemp < 90 {
					nonMatchStringQuestion = qnas[i].Question
					println(nonMatchStringQuestion)
					nonMatchStringAnswer = qnas[i].Answer
					println(nonMatchStringAnswer)
					similarityNonMatch = float64(similarityTemp)
					newTuple := stringSimilarity{nonMatchStringQuestion, nonMatchStringAnswer, similarityNonMatch}
					nonMatchStrings = append(nonMatchStrings, newTuple)
				}
				if index != -1 && similarityTemp >= 90 && similarityTemp > float32(similarityMatch) {
					matchAnswerString = qnas[i].Answer
					isMatch = true
				}
			} else {
				// bm
				inputTemp := strings.ReplaceAll(input.Input, " ", "")

				qnaTemp := strings.ReplaceAll(qnas[i].Question, " ", "")
				index, similarityTemp := algorithm.BMMatch(inputTemp, qnaTemp)

				if index == -1 && similarityTemp < 90 {
					nonMatchStringQuestion = qnas[i].Question
					nonMatchStringAnswer = qnas[i].Answer
					similarityNonMatch = float64(similarityTemp)
					newTuple := stringSimilarity{nonMatchStringQuestion, nonMatchStringAnswer, similarityNonMatch}
					nonMatchStrings = append(nonMatchStrings, newTuple)
				}
				if index != -1 && similarityTemp >= 90 && similarityTemp > float32(similarityMatch) {
					matchAnswerString = qnas[i].Answer
					isMatch = true
				}
			}
		}
		if isMatch {
			answer = matchAnswerString

		} else {
			sort.Slice(nonMatchStrings, func(i, j int) bool {
				return nonMatchStrings[i].similarityPercentage > nonMatchStrings[j].similarityPercentage
			})
			answer += "Pertanyaan tidak ditemukan di database.\n"
			answer += "Apakah maksud anda: \n"
			for i := 1; i <= 3; i++ {
				answer += (string(i) + ". " + nonMatchStrings[i].nonMatchStringQuestion + "\n")
			}

		}

	}
	newInput := model.InputUser{InputText: input.Input, Algorithm: input.Algorithm, Answer: answer}
	if err := database.DB.Create(&newInput).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error!!",
		})
	}
	return c.JSON(newInput)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("question")

	var input model.InputUser
	if database.DB.Delete(&input, id).RowsAffected == 0 {
		println("masuk ke sini dia ges")
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Can not delete the data",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Data is successfully deleted!",
	})
}
