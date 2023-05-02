package algorithm

import (
	"regexp"
	"strconv"
	"math"
	"strings"
)

func calculateMathOperation(input string) string {
	matharray := regexp.MustCompile(`(\d+|\(|\)|\^\s*|\*\s*|/\s*|\+\s*|-)`)
	matcharray := matharray.FindAllStringSubmatch(input, -1)

	buffer := make([]string, len(matcharray))
	counter := 0
	for _, match := range matcharray {
		if match[1] != " " {
			buffer[counter] = match[1]
			counter++
		}
	}

	var result float64

	for i := 0; i < len(buffer); i++ {
		if buffer[i] == "(" {
			parentCount := 1 
			j := i + 1
			for ; j < len(buffer); j++ {
				if (buffer[j] == "(") {
					parentCount++
				} else if (buffer[j] == ")") {
					parentCount--
				}
				if parentCount == 0 {
					break
				}
			}
			subExpr := strings.Join(buffer[i+1:j], "")
			subResult := calculateMathOperation(subExpr)
			buffer[i] = subResult
			buffer = append(buffer[:i+1], buffer[j+1:]...)
			i--
		}	
	}


	if len(buffer) == 3 {
		operand1, _ := strconv.ParseFloat(buffer[0], 64)
		operator := buffer[1]
		operand2, _ := strconv.ParseFloat(buffer[2], 64)
		switch operator {
		case "+":
			result = operand1 + operand2
			return strconv.FormatFloat(result, 'f', 2, 64)
		case "-":
			result = operand1 - operand2
			return strconv.FormatFloat(result, 'f', 2, 64)
		case "*":
			result = operand1 * operand2
			return strconv.FormatFloat(result, 'f', 2, 64)
		case "/":
			result = operand1 / operand2
			return strconv.FormatFloat(result, 'f', 2, 64)
		case "^":
			result = math.Pow(operand1, operand2)
			return strconv.FormatFloat(result, 'f', 2, 64) 
		}
	} else {
		i := 0
		operand1, _ := strconv.ParseFloat(buffer[0], 64)
		operator := buffer[1]
		operand2, _ := strconv.ParseFloat(buffer[2], 64)
		if operator == "^" {
			result = math.Pow(operand1, operand2)
			buffer = removeElement(buffer, 0)
			buffer = removeElement(buffer, 0)
			buffer = removeElement(buffer, 0)
			e := strconv.FormatFloat(result, 'f', 2, 64)
			buffer = insertElement(buffer, e, 0)
		} else if operator == "+"  || operator == "-" || operator == "*" || operator == "/" {
			i += 2
		}

		// check for additional operation
		for i < len(buffer) - 2 {
			operand3, _ := strconv.ParseFloat(buffer[i], 64)
			op := buffer[i+1]
			operand4, _ := strconv.ParseFloat(buffer[i+2], 64)
			if op == "^" {
				result = math.Pow(operand3, operand4)
				buffer = removeElement(buffer, i)
				buffer = removeElement(buffer, i)
				buffer = removeElement(buffer, i)
				e := strconv.FormatFloat(result, 'f', 2, 64)
				buffer = insertElement(buffer, e, i)
			} else if op == "+" || op == "-" || op == "*" || op == "/" {
				i += 2
			} else {
				i += 2
			}
			
		}
	}


	if len(buffer) >= 3 {
		i := 0
		operand1, _ := strconv.ParseFloat(buffer[0], 64)
		operator := buffer[1]
		operand2, _ := strconv.ParseFloat(buffer[2], 64)
		if operator == "*" {
			result = operand1 * operand2
			buffer = removeElement(buffer, 0)
			buffer = removeElement(buffer, 0)
			buffer = removeElement(buffer, 0)
			e := strconv.FormatFloat(result, 'f', 2, 64)
			buffer = insertElement(buffer, e, 0)
		} else if operator == "/" {
			result = operand1 * operand2
			buffer = removeElement(buffer, 0)
			buffer = removeElement(buffer, 0)
			buffer = removeElement(buffer, 0)
			e := strconv.FormatFloat(result, 'f', 2, 64)
			buffer = insertElement(buffer, e, 0)
		} else if operator == "+"  || operator == "-" {
			i += 2
		}
		
		// check for additional operation
		for i < len(buffer) - 2 {
			operand3, _ := strconv.ParseFloat(buffer[i], 64)
			op := buffer[i+1]
			operand4, _ := strconv.ParseFloat(buffer[i+2], 64)
			if op == "*" {
				result = operand3 * operand4
				buffer = removeElement(buffer, i) 
				buffer = removeElement(buffer, i) 
				buffer = removeElement(buffer, i)
				e := strconv.FormatFloat(result, 'f', 2, 64)
				buffer = insertElement(buffer, e, i)
			} else if op == "/" {
				result = operand3 / operand4
				buffer = removeElement(buffer, i) 
				buffer = removeElement(buffer, i)
				buffer = removeElement(buffer, i)
				e := strconv.FormatFloat(result, 'f', 2, 64)
				buffer = insertElement(buffer, e, i)
			} else if op == "+" {
				i += 2
			} else {
				i += 2
			}
			
		}
	}

	// check for remaining addition and substraction operation
	if len(buffer) >= 3 {
		operand1, _ := strconv.ParseFloat(buffer[0], 64)
		operator := buffer[1]
		operand2, _ := strconv.ParseFloat(buffer[2], 64)

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
		for i := 3; i < len(buffer); i+=2 {
			op := buffer[i]
			operand3, _ := strconv.ParseFloat(buffer[i+1], 64)
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


	return strconv.FormatFloat(result, 'f', 2, 64)
}

func removeElement(arr []string, index int) []string {
    result := make([]string, len(arr)-1, cap(arr))
    copy(result, arr[:index])
    copy(result[index:], arr[index+1:])
    return result
}


func insertElement(arr []string, element string, index int) []string {
    result := make([]string, len(arr)+1)
    copy(result, arr[:index])
    result[index] = element
    copy(result[index+1:], arr[index:])
    return result
}
