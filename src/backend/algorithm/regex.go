package algorithm

import (
	"regexp"
)

func isMathOperation(pattern string) bool {
	operation := `^(\d+)\s*([\-\+\*\/])\s*(\d+)(\s*([\-\+\*\/])\s*(\d+)){0,}.*$`
	regex := regexp.MustCompile(operation)
	return regex.MatchString(pattern)
}

func isDate(pattern string) bool {
	regex := regexp.MustCompile(`^([Hh][aA][rR][iI]\s*[aA][Pp][Aa]\s*)?\d{2}/\d{2}/\d{4}.*$`)
	return regex.MatchString(pattern)
}

func isAddingQNAToDatabase(pattern string) bool {
	regex := regexp.MustCompile(`[tT][aA][mM][bB][aA][hH][kK][aA][nN]\s*[pP][eE][Rr][tT][aA][nN][yY][aA][aA][nN]\s*.*
	[dD][eE][nN][gG][aA][nN]\s*[jJ][aA][wW][aA][bB][aA][nN]\s*.*$`)
	return regex.MatchString(pattern)
}
	
func isErasingQuestion(pattern string) bool {
	regex := regexp.MustCompile(`[hH][aA][pP][uU][sS]\s*[pP][eE][rR][tT][aA][nN][yY][aA][aA][nN]\s*.*$`)
	return regex.MatchString(pattern)
}

func checkQuestion(input string) string {
	var ans string = ""
	if isMathOperation(input) {
		ans = "mathoperation"
	} 
	if isDate(input) {
		dateparse, _ := regexp.Compile(`(\d{2})/(\d{2}/)\(d{4})`)
		date := dateparse.FindStringSubmatch(input)
		day := calculateDate(date) 
		ans = day
	} 
	if isAddingQNAToDatabase(input) {
		ans = "adding"
	} 
	if isErasingQuestion(input) {
		ans = "erasing"
	}
	return ans
}


