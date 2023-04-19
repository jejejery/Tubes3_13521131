package algorithm

import "regexp"

func isMathOperation(pattern string) bool {
	regex := regexp.MustCompile(`^\s*\d+(\s*[-+*/]\s*\d+)+\s*$`)
	return regex.MatchString(pattern)
}

func isDate(pattern string) bool {
	regex := regexp.MustCompile(`^\d{2}/\d{2}/\d{4}$`)
	return regex.MatchString(pattern)
}

func isAddingQNAToDatabase(pattern string) bool {
	regex := regexp.MustCompile(`[tT][aA][mM][bB][aA][hH][kK][aA][nN]$`)
	return regex.MatchString(pattern)
}
	
func isErasingQuestion(pattern string) bool {
	regex := regexp.MustCompile(`[hH][aA][pP][uS]`)
	return regex.MatchString(pattern)
}

func checkQuestion(input string) string {
	var ans string
	if isMathOperation(input) {
		ans = "mathoperation"
	} 
	if isDate(input) {
		ans = "date"
	} 
	if isAddingQNAToDatabase(input) {
		ans = "adding"
	} 
	if isErasingQuestion(input) {
		ans = "erasing"
	}
	return ans
}
