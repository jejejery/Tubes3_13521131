package inputcontroller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"regexp"
	"sort"
	"strconv"
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
		Session   int64  `json:"Session"`
		Input     string `json:"Input"`
		Algorithm bool   `json:"Algorithm"`
	}

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request!!",
		})
	}

	var answer string     // answer to return
	var ansArray []string // temp answer from checkQuestion
	ansArray = algorithm.CheckQuestion(input.Input, ansArray)
	for i := 0; i < len(ansArray); i++ {
		println(ansArray[i])
	}
	println("cek")

	// database.DB.Find(&qnas)
	for i := 0; i < len(ansArray); i++ {
		_, err := strconv.ParseFloat(ansArray[i], 64)
		if err == nil || ansArray[i] == "Sintaks persamaan tidak valid!" {
			answer += ansArray[i]
		} else if algorithm.IsDayOutput(ansArray[i]) {
			answer += ansArray[i]
		} else if algorithm.IsAddingQNAToDatabase(ansArray[i]) {
			var newQnA model.QnA
			pattern := `(?i)Tambahkan pertanyaan (.+) dengan jawaban (.+)`
			regex := regexp.MustCompile(pattern)
			matches := regex.FindStringSubmatch(ansArray[i])
			newQnA.Question = matches[1]
			newQnA.Answer = matches[2]
			jsonStr, err := json.Marshal(newQnA)
			if err != nil {
				// return err
			}
			req, err := http.NewRequest("POST", "http://localhost:8000/api/qna", bytes.NewBuffer((jsonStr)))
			if err != nil {
				return err
			}
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			// var qna
		} else if algorithm.IsErasingQuestion((ansArray[i])) {
			var deleteQnA model.QnA
			pattern := `(?i)hapus pertanyaan (.+)`
			regex := regexp.MustCompile(pattern)
			matches := regex.FindStringSubmatch(ansArray[i])
			deleteQnA.Question = matches[1]
			deleteQnA.Answer = ""
			jsonStr, err := json.Marshal(deleteQnA)
			if err != nil {
				return err
			}
			req, err := http.NewRequest("DELETE", "http://localhost:8000/api/qna", bytes.NewBuffer((jsonStr)))
			if err != nil {
				return err
			}
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

		} else {
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
			var qnas []model.QnA
			database.DB.Find(&qnas)
			var nonMatchStrings []stringSimilarity
			n := len(qnas)
			println(n)
			println(ansArray[i])
			for j := 0; j < n; j++ {
				if input.Algorithm {
					// kmp
					var x string = ansArray[i]
					var y string = qnas[j].Question
					index, similarityTemp := algorithm.KMPMatch(x, y)
					println(index)
					println(similarityTemp)

					if index == -1 {
						nonMatchStringQuestion = qnas[j].Question
						nonMatchStringAnswer = qnas[j].Answer
						similarityNonMatch = float64(similarityTemp)
						newTuple := stringSimilarity{nonMatchStringQuestion, nonMatchStringAnswer, similarityNonMatch}
						nonMatchStrings = append(nonMatchStrings, newTuple)
					}
					if index != -1 && similarityTemp >= 90 && similarityTemp > float32(similarityMatch) {
						matchAnswerString = qnas[j].Answer
						isMatch = true
					}
				} else {
					// bm
					inputTemp := strings.ReplaceAll(input.Input, " ", "")

					qnaTemp := strings.ReplaceAll(qnas[j].Question, " ", "")
					index, similarityTemp := algorithm.BMMatch(inputTemp, qnaTemp)

					if index == -1 {
						nonMatchStringQuestion = qnas[j].Question
						nonMatchStringAnswer = qnas[j].Answer
						similarityNonMatch = float64(similarityTemp)
						newTuple := stringSimilarity{nonMatchStringQuestion, nonMatchStringAnswer, similarityNonMatch}
						nonMatchStrings = append(nonMatchStrings, newTuple)
					}
					if index != -1 && similarityTemp >= 90 && similarityTemp > float32(similarityMatch) {
						matchAnswerString = qnas[j].Answer
						isMatch = true
					}
				}
			}
			if isMatch {
				answer += matchAnswerString

			} else {
				sort.Slice(nonMatchStrings, func(i, j int) bool {
					return nonMatchStrings[i].similarityPercentage > nonMatchStrings[j].similarityPercentage
				})
				if nonMatchStrings[0].similarityPercentage > 90 {
					answer += nonMatchStrings[0].nonMatchStringAnswer
				} else {
					answer += "Pertanyaan tidak ditemukan di database.\n"
					answer += "Apakah maksud anda: \n"
					for i := 1; i <= 3; i++ {
						answer += (strconv.Itoa(i) + ". " + nonMatchStrings[i-1].nonMatchStringQuestion + "\n")

					}
				}
			}
		}
	}
	newInput := model.InputUser{Session: input.Session, InputText: input.Input, Algorithm: input.Algorithm, Answer: answer}
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
