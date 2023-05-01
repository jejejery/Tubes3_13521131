package algorithm

import (
	"regexp"
	"strconv"
)

func MathOperation(pattern string) bool {
	operation := `^(-?\d+)\s*([-+*\/])\s*([-+*\/])?\s*(-?\d+)(\s*([-+*\/])\s*([-+*\/])?\s*(-?\d+)){0,}.*`
	regex := regexp.MustCompile(operation)
	return regex.MatchString(pattern)
}

func isMathOperationValid(pattern string) bool {
	operation := `^(-?\d+)\s*([\-\+\*\/])\s*(-?\d+)(\s*([\-\+\*\/])\s*(-?\d+)){0,}.*$`
	regex := regexp.MustCompile(operation)
	return regex.MatchString(pattern)
}

func isDate(pattern string) bool {
	regex := regexp.MustCompile(`^([Hh][aA][rR][iI]\s*[aA][Pp][Aa]\s*)?\d{2}/\d{2}/\d{4}.*$`)
	return regex.MatchString(pattern)
}

func isAddingQNAToDatabase(pattern string) bool {
	regex := regexp.MustCompile(`[tT][aA][mM][bB][aA][hH][kK][aA][nN]\s*[pP][eE][rR][tT][aA][nN][yY][aA][aA][nN]\s*.*\s*[dD][eE][nN][gG][aA][nN]\s*[jJ][aA][wW][aA][bB][aA][nN]\s*.*\s*.*`)
	return regex.MatchString(pattern)
}

func isErasingQuestion(pattern string) bool {
	regex := regexp.MustCompile(`[hH][aA][pP][uU][sS]\s*[pP][eE][rR][tT][aA][nN][yY][aA][aA][nN]\s*.*$`)
	return regex.MatchString(pattern)
}

func CheckQuestion(input string) string {
	var ans string = ""
	if MathOperation(input) {
		if isMathOperationValid(input) {
			ans = calculateMathOperation(input)
		} else {
			ans = "Sintaks persamaan tidak valid!"
		}
	}
	if isDate(input) {
		dateparse, _ := regexp.Compile(`(\d{2})/(\d{2})/(\d{4})`)
		date := dateparse.FindString(input)
		dayStr := string(date[0:2])
		monthStr := string(date[3:5])
		yearStr := string(date[6:10])
		day, _ := strconv.Atoi(dayStr)
		month, _ := strconv.Atoi(monthStr)
		year, _ := strconv.Atoi(yearStr)
		if day < 0 || month < 0 || year < 0 {
			ans = "Masukan tanggal tidak valid!"
		} else if month == 2 && day > 29 {
			ans = "Masukan tanggal tidak valid!"
		} else {
			day := calculateDate(date)
			ans = day
		}
	}
	if isAddingQNAToDatabase(input) {
		ans = "Question is added to database!"
	}
	if isErasingQuestion(input) {
		ans = "Question is deleted from database!"
	}
	return ans
}
