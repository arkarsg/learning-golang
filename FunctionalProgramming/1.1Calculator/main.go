package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type alias for functions
type operateFunc func(int, int) int

// define basic operations
func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func div(a, b int) int {
	if b == 0 {
		panic("divide by zero")
	}
	return a / b
}

var operationsMap = map[string]operateFunc{
	"+":  add,
	"-":  subtract,
	"*":  mult,
	"/":  div,
	"<<": func(a, b int) int { return a << b },
}

func doOp(op string, a int, b int) int {
	if f, ok := operationsMap[op]; ok {
		return f(a, b)
	} else {
		panic("Op not supported")
	}
}

func main() {
	fmt.Println(
		"Simple calculator with dispatcher pattern\n" +
			"Supported operations: + - * /\n" +
			"Format: [num1] [op] [num2]",
	)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter expression: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Fields(input)
	if len(parts) != 3 {
		fmt.Println("Invalid format. Please enter in the format: [num1] [op] [num2]")
		return
	}

	// Parse the numbers and the operator
	num1, err1 := strconv.Atoi(parts[0])
	op := parts[1]
	num2, err2 := strconv.Atoi(parts[2])

	if err1 != nil || err2 != nil {
		fmt.Println("Invalid numbers. Please enter valid integers.")
		return
	}

	// Perform the calculation based on the operator
	result := doOp(op, num1, num2)
	fmt.Println(result)
}
