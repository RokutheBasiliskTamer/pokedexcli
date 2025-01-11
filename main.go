package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	var cleanedInput []string
	words := strings.Fields(strings.ToLower(text))

	cleanedInput = append(cleanedInput, words...)

	return cleanedInput
}

func main() {
	fmt.Printf("Hello, World!\n")
}
