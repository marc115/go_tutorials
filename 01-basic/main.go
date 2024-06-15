package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 12
	fmt.Println(intNum)

	var floatNum float64 = .0
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2

	var result = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)

	fmt.Println(utf8.RuneCountInString("a")) // get number of characters in strings

	var a rune = 'a' // a single character
	fmt.Println(a)

	var myBoolean bool = true
	fmt.Println(myBoolean)

	myBoolean2 := true //shorthand way to create a variable
	fmt.Print(myBoolean2)

	var var1, var2 int = 1, 2 //defining multiple variables
	fmt.Println(var1, var2)

	const pi float32 = 3.14
}
