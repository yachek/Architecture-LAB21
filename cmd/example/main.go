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
	//scanner := bufio.NewScanner(os.Stdin)
	//scanner.Scan()
	//input = scanner.Text()
	res, _ := go21.PostfixToInfix(input)
	fmt.Println(res)
}
