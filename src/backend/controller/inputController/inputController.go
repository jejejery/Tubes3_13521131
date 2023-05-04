package inputcontroller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/jejejery/src/backend/algorithm"
	database "github.com/jejejery/src/backend/db"
	"github.com/jejejery/src/backend/model"
)

func Index(c *fiber.Ctx) error {
	var input []model.InputUser
	database.DB.Find(&input)
	return c.JSON(input)
}

func Show(c *fiber.Ctx) error {

	session, err := strconv.ParseInt(c.Query("Session"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request!!",
		})
	}
	var interactions []model.InputUser

	database.DB.Where("Session = ?", session).Find(&interactions)
	return c.JSON(interactions)
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
		_, err := strconv.ParseFloat(ansArray[i], 64)
		if err == nil || ansArray[i] == "Sintaks persamaan tidak valid!" {
			answer += ansArray[i] + "\n"
		} else if err == nil || ansArray[i] == "Masukan tanggal tidak valid!" {
			answer += ansArray[i] + "\n"		
		} else if algorithm.IsDayOutput(ansArray[i]) {
			answer += ansArray[i] + "\n"
		} else if algorithm.IsAddingQNAToDatabase(ansArray[i]) {
			var newQnA model.QnA
			pattern := `(?i)Tambahkan pertanyaan (.+) dengan jawaban (.+)`
			regex := regexp.MustCompile(pattern)
			matches := regex.FindStringSubmatch(ansArray[i])
			newQnA.Question = matches[1]
			newQnA.Answer = matches[2]
			jsonStr, err := json.Marshal(newQnA)
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
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			answer += string(bodyBytes)

			// var qna
		} else if algorithm.IsErasingQuestion((ansArray[i])) {
			var request struct {
				Question  string
				Answer    string
				Algorithm bool
			}
			pattern := `(?i)hapus pertanyaan (.+)`
			regex := regexp.MustCompile(pattern)
			matches := regex.FindStringSubmatch(ansArray[i])
			request.Question = matches[1]
			request.Answer = ""
			request.Algorithm = input.Algorithm

			jsonStr, err := json.Marshal(request)
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
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			answer += string(bodyBytes)

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
