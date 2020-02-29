package go21

import (
	"fmt"
	"strings"
)

const (
	operators = "+-*/^"
	digits    = "1234567890"
)

//PostfixToInfix converts math statement
func PostfixToInfix(input string) (string, error) {
	var result string = ""
	if input == "" {
		return "", fmt.Errorf("Input error: the input expression shall not be empty")
	}
	var arr = strings.Split(input, " ")
	err := testArray(arr)
	if err != nil {
		return "", err
	}
	result += arr[0]
	for i := 1; i < len(arr); i += 2 {
		if !strings.ContainsAny(arr[i], digits) ||
			!strings.ContainsAny(arr[i+1], operators) {
			return "", fmt.Errorf("Input error: the input expression is not correct as a postfix expression")
		}

		result += " " + arr[i+1] + " " + arr[i]
		if bracesNeeded(arr, i+1) {
			result = "(" + result + ")"
		}
	}
	return result, nil
}

func bracesNeeded(arr []string, i int) bool {
	if i >= len(arr)-1 {
		return false
	}
	if strings.ContainsAny(arr[i], "+-") &&
		strings.ContainsAny(arr[i+2], "*/") {
		return true
	} else if strings.ContainsAny(arr[i], "*/") &&
		strings.ContainsAny(arr[i+2], "^") {
		return true
	}
	return false
}

func testArray(arr []string) error {
	if len(arr)%2 == 0 {
		return fmt.Errorf("Input error: the input expression is not correct as a postfix expression")
	}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if !strings.ContainsAny(string(arr[i][j]), operators+digits) {
				return fmt.Errorf("Input error: the only allowed characters in input exprression are digits and math operators")
			}
		}
		if strings.ContainsAny(arr[i], operators) &&
			strings.ContainsAny(arr[i], digits) {
			return fmt.Errorf("Input error: every item in input expression must be separated by space character")
		}

	}
	return nil
}
