package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Println("The result of the division is %v", result)
	// } else {
	// 	fmt.Println("The result of the division is %v with remainder %v", result, remainder)
	// }
	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Println("The result of the integer division is %v", result)
	default:
		fmt.Println("The result of the integer division is %v with remainder %v", result, remainder)
	}
	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	return numerator / denominator, numerator % denominator, err
}
