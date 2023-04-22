package algorithm

import (
	"regexp"
	"strconv"
)

func calculateMathOperation(input string) string {
	// matharray := regexp.MustCompile(`(-?\d+)\s*([-+*\/])\s*(-?\d+)((?:\s*([-+*\/])\s*(-?\d+))+)`)
	matharray := regexp.MustCompile(`(\d+)\s{0,}([-+*\/]|)`)
	
	matcharray := matharray.FindAllStringSubmatch(input, -1)
	
	buffer := make([]string, 2*(len(matcharray)-1) + 1)
	counter := 0
	counterMatch := 0
	for _, match := range matcharray {
		if (counterMatch == len(matcharray)-1) {
			buffer[counter] = match[1]
 		} else {
			if len(match) >= 3 {
				for i := 1; i < len(match); i++ {
					buffer[counter] = match[i]
					counter++
				}
				counterMatch++
			}
		}
	}
	
	var result int 

	// evaluate the multiplication and division first
		
	if len(buffer) == 3 {
		operand1, _ := strconv.Atoi(buffer[0])
		operator := buffer[1]
		operand2, _ := strconv.Atoi(buffer[2])
		switch operator {
		case "+":
			result = operand1 + operand2
			return strconv.Itoa(result)
		case "-":
			result = operand1 - operand2
			return strconv.Itoa(result)
		case "*":
			result = operand1 * operand2
			return strconv.Itoa(result)
		case "/":
			result = operand1 / operand2
			return strconv.Itoa(result)
		}
	} else {
		i := 0
		operand1, _ := strconv.Atoi(buffer[0])
		operator := buffer[1]
		operand2, _ := strconv.Atoi(buffer[2])
		if operator == "*" {
			result = operand1 * operand2
			buffer = removeElement(buffer, 0)
			buffer = removeElement(buffer, 0)
			buffer = removeElement(buffer, 0)
			e := strconv.Itoa(result)
			buffer = insertElement(buffer, e, 0)
		} else if operator == "/" {
			result = operand1 * operand2
			buffer = removeElement(buffer, 0)
			buffer = removeElement(buffer, 0)
			buffer = removeElement(buffer, 0)
			e := strconv.Itoa(result)
			buffer = insertElement(buffer, e, 0)
		} else if operator == "+"  || operator == "-" {
			i += 2
		}
		
		// check for additional operation
		for i < len(buffer) - 2 {
			operand3, _ := strconv.Atoi(buffer[i])
			op := buffer[i+1]
			operand4, _ := strconv.Atoi(buffer[i+2])
			if op == "*" {
				result = operand3 * operand4
				buffer = removeElement(buffer, i) 
				buffer = removeElement(buffer, i) 
				buffer = removeElement(buffer, i)
				e := strconv.Itoa(result)
				buffer = insertElement(buffer, e, i)
			} else if op == "/" {
				result = operand3 / operand4
				buffer = removeElement(buffer, i) 
				buffer = removeElement(buffer, i)
				buffer = removeElement(buffer, i)
				e := strconv.Itoa(result)
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
		operand1, _ := strconv.Atoi(buffer[0])
		operator := buffer[1]
		operand2, _ := strconv.Atoi(buffer[2])

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
			operand3, _ := strconv.Atoi(buffer[i+1])
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
