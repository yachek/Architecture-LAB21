package go21

import (
	"strings"
)

//PostfixToInfix converts math statement
func PostfixToInfix(input string) (string, error) {
	var result string = ""
	var arr = strings.Split(input, " ")
	result += arr[0]
	for i := 1; i < len(arr); i += 2 {
		result += " " + arr[i+1] + " " + arr[i]
		if bracesNeeded(arr, i+1) {
			result = "(" + result + ")"
		}
	}
	return result, nil //fmt.Errorf("TODO")
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
