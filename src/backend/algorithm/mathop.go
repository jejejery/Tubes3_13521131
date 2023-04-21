package algorithm

import (
	"regexp"
	"strconv"
	"fmt"
)

func calculateMathOperation(input string) string {
	matharray := regexp.MustCompile(`(\d+)\s*([-+*/])\s*(\d+)(?:\s*([-+*/])\s*(\d+))*`)
	
	matcharray := matharray.FindAllStringSubmatch(input, -1) 

	fmt.Println(matcharray)

	result, _ := strconv.Atoi(matcharray[0][1])
	for i := 0; i < len(matcharray[0])-1; i+=2 {
		operator := matcharray[0][i+2]
		operand, _ := strconv.Atoi(matcharray[0][i+3])
		switch operator {
		case "+":
			result += operand
		case "-":
			result -= operand
		case "*":
			result *= operand
		case "/":
			result /= operand
		}

		fmt.Println(operator)
		fmt.Println(i)
		fmt.Println(len(matcharray[0]))
		if i < len(matcharray[0])-3 {
			nextOperator := matcharray[0][i+4]
			if nextOperator == "*" || nextOperator == "/" {
				continue
			}
		}
	}

	return strconv.Itoa(result)
}