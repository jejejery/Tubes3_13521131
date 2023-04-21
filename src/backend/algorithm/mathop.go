package algorithm

import (
	"regexp"
	"strconv"
)

func calculateMathOperation(input string) string {
	matharray := regexp.MustCompile(`(\d+)\s*([-+*/])\s*(\d+)(?:\s*([-+*/])\s*(\d+))*`)
	
	matcharray := matharray.FindAllStringSubmatch(input, -1) 

	var result int 
	for _, match := range matcharray {
		operand1, _ := strconv.Atoi(match[1])
		operator := match[2]
		operand2, _ := strconv.Atoi(match[3])

		switch operator {
		case "+":
			result = operand1 + operand2
		case "-":
			result = operand1 - operand2
		case "*":
			result = operand1 * operand2
		case "/":
			result = operand1 / operand2
		}

		// check for additional operation
		for i := 4; i < len(match); i+=2 {
			op := match[i]
			operand3, _ := strconv.Atoi(match[i+1])
			switch op {
			case "+":
				result += operand3
			case "-":
				result -= operand3
			case "*":
				result *= operand3
			case "/":
				result /= operand3
			}
			
		}
	}

	return strconv.Itoa(result)
}