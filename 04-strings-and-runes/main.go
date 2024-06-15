package main

import (
	"fmt"
	"strings"
)

func main() {
	for pos, char := range "日本語" { //iterating through a string
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}

	var strSlice = []string{"h", "e", "l", "l", "o"}
	var strBuilder strings.Builder //strings are inmutable in go. to dynamically make a string we can use a stringBuilder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}

	var catStr = strBuilder.String()
	fmt.Println("%v", catStr)
}
