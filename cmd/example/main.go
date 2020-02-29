package main

import (
	go21 "GO21"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Build version: " + version)
	input := strings.Join(os.Args[1:], " ")
	res, err := go21.PostfixToInfix(input)
	if err != nil {
		fmt.Println(rr)
	}
	fmt.Println(res)
}
