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

// func isQuestionInDataBase(pattern string) bool {
// 	regex := regex.MustCompile(//regex expression for question that is stored in database*/)
// 	return regex.MatchString(pattern)
// }

// func isAddingQNAToDatabase(pattern string) bool {
	// 	regex := regex.MustCompile(//regex expression for tambah pertanyaan)
	// 	return regex.MatchString(pattern)
	// }
	
// func isErasingQuestion(pattern string) bool {
	// 	regex := regex.MustCompile(//regex expression for menghapus pertanyaan)
	// 	return regex.MatchString(pattern)
// }