package main

import (
	"errors"
	"fmt"
)

func printMe(printValue string) {
	fmt.Println(printValue)

}
func printNum(a int, b int) {
	fmt.Println(a, b)
}
func product(a int, b int) int {
	return a * b
}
func addNum(a int, b int) int {
	return a + b
}
func divNum(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("division by zero")
		return 0, 0, err

	}
	var result int = a / b
	var remainder int = a % b
	return result, remainder, err

}
func main() {
	printValue := "Yes i have done it"
	printMe(printValue)
	a := 2
	b := 0
	printNum(a, b)
	fmt.Println(addNum(a, b))
	fmt.Println(product(a, b))

	result, remainder, err := divNum(a, b)
	if err != nil {
		fmt.Println("Error:", err)

	} else if remainder == 0 {
		fmt.Println("Result:", result)

	} else {
		fmt.Println("Print the result of integer division %v with remainder %v", result, remainder)

	}

}
