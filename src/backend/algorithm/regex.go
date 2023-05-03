package algorithm

import (
	"regexp"
	"strconv"
	"strings"
)

func MathOperation(pattern string) bool {
	operation := `^\(?\s*\(?\s*(-?\d+)\s*([\^\-+*\/])\s*([\^\-+*\/])?\s*(-?\d+)\)?\s*(\s*([\^\-+*\/])\s*([\^\-+*\/])?\s*(-?\d+)){0,}\)?\s*.*(\s.*)*(\n.*)*$`
	regex := regexp.MustCompile(operation)
	return regex.MatchString(pattern)
}

func IsMathOperationValid(pattern string) bool {
	operation := `^\(?\s*\(?\s*(-?\d+)\s*([\^\-\+\*\/])\s*(-?\d+)\)?\s*(\s*([\^\-\+\*\/])\s*(-?\d+)){0,}\)?.*(\s.*)*(\n.*)*$`
	regex := regexp.MustCompile(operation)
	if !regex.MatchString(pattern) {
		return false
	}
	// check the parentheses
	stack := []rune{}
	for _, char := range pattern {
		if char == '(' {
			stack = append(stack, char)
		} else if char == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
                return false
            }
            stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func IsDate(pattern string) bool {
	regex := regexp.MustCompile(`^([Hh][aA][rR][iI]\s*[aA][Pp][Aa]\s*)?\d{2}\/\d{2}\/\d{4}(\s.*)*(\n.*)*$`)
	return regex.MatchString(pattern)
}

func IsAddingQNAToDatabase(pattern string) bool {
	regex := regexp.MustCompile(`^[tT][aA][mM][bB][aA][hH][kK][aA][nN]\s*[pP][eE][rR][tT][aA][nN][yY][aA][aA][nN]\s*.*\s*[dD][eE][nN][gG][aA][nN]\s*[jJ][aA][wW][aA][bB][aA][nN].*(\s.*)*(\n.*)*$`)
	return regex.MatchString(pattern)
}

func IsErasingQuestion(pattern string) bool {
	regex := regexp.MustCompile(`^[hH][aA][pP][uU][sS]\s*[pP][eE][rR][tT][aA][nN][yY][aA][aA][nN].*(\s.*)*(\n.*)*$`)
	return regex.MatchString(pattern)
}
func IsDayOutput(pattern string) bool {
	day := [8]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu", "Masukan tanggal tidak valid!"}
	for i := 0; i < 8; i++ {
		if pattern == day[i] {
			return true
		}
	}
	return false
}

// manggil check question
// var ansArray []string
// chechQuestion(input, ansArray)

func CheckQuestion(input string, ansArray []string) []string {
	if len(input) == 0 {
		return ansArray
	}

	var ans string
	if IsDate(input) {
		dateparse, _ := regexp.Compile(`(\d{2})/(\d{2})/(\d{4})`)
		date := dateparse.FindString(input)
		dayStr := string(date[0:2])
		monthStr := string(date[3:5])
		day, _ := strconv.Atoi(dayStr)
		month, _ := strconv.Atoi(monthStr)
		if month == 2 && day > 29 {
			ans = "Masukan tanggal tidak valid!"
			ansArray = append(ansArray, ans)
		} else {
			day := calculateDate(date)
			ans = day
			ansArray = append(ansArray, ans)
		}
		pattern := `([Hh][aA][rR][iI]\s*[aA][Pp][Aa]\s*)?\d{2}/\d{2}/\d{4}\s*\n*`
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(input)
		input = strings.Replace(input, matches[0], "", 1)
	} else if MathOperation(input) {
		if IsMathOperationValid(input) {
			pattern := `^\(?\(?(-?\d+)\s*([\^\-\+\*\/])\s*(-?\d+)\)?(\s*([\^\-\+\*\/])\s*(-?\d+)){0,}\)?.*\s*\n*$`
			re := regexp.MustCompile(pattern)
			matches := re.FindStringSubmatch(input)
			ans = calculateMathOperation(matches[0])
			ansArray = append(ansArray, ans)
		} else {
			ans = "Sintaks persamaan tidak valid!"
			ansArray = append(ansArray, ans)
		}
		pattern := `^\(?\s*\(?\s*(-?\d+)\s*([\^\-+*\/])\s*([\^\-+*\/])?\s*(-?\d+)\)?\s*(\s*([\^\-+*\/])\s*([\^\-+*\/])?\s*(-?\d+)){0,}\)?.*\s*\n*$`
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(input)
		input = strings.Replace(input, matches[0], "", 1)
	} else if IsAddingQNAToDatabase(input) {
		pattern := `[tT][aA][mM][bB][aA][hH][kK][aA][nN]\s*[pP][eE][rR][tT][aA][nN][yY][aA][aA][nN]\s*.*\s*[dD][eE][nN][gG][aA][nN]\s*[jJ][aA][wW][aA][bB][aA][nN].*\n*`
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(input)
		matches_clean := strings.Replace(matches[0], "\n", "", -1)
		if strings.HasSuffix(matches_clean, " ") {
			matches_clean = strings.TrimRight(matches_clean, " ")
		}
		ansArray = append(ansArray, matches_clean)
		if re.MatchString(input) {
			input = re.ReplaceAllString(input, "$1")
		} else {
			input = ""
		}
	} else if IsErasingQuestion(input) {
		pattern := `[hH][aA][pP][uU][sS]\s*[pP][eE][rR][tT][aA][nN][yY][aA][aA][nN].*\n*$`
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(input)
		matches_clean := strings.Replace(matches[0], "\n", "", -1)
		if strings.HasSuffix(matches_clean, " ") {
			matches_clean = strings.TrimRight(matches_clean, " ")
		}
		ansArray = append(ansArray, matches_clean)
		if re.MatchString(input) {
			input = re.ReplaceAllString(input, "$1")
		} else {
			input = ""
		}
	} else {
		pattern := `.*\s*\n*`
		re := regexp.MustCompile(pattern)
		matches := re.FindStringSubmatch(input)
		matches_clean := strings.Replace(matches[0], "\n", "", -1)
		if strings.HasSuffix(matches_clean, " ") {
			matches_clean = strings.TrimRight(matches_clean, " ")
		}
		ansArray = append(ansArray, matches_clean)
		input = strings.Replace(input, matches[0], "", 1)
		return CheckQuestion(input, ansArray)
	}
	return CheckQuestion(input, ansArray)
}
