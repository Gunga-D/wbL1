package main

import (
	"fmt"
	"strings"
)

func InvertStringsWords(input string) string {
	var result []string

	words := strings.Fields(input)
	for idx := len(words) - 1; idx >= 0; idx-- {
		result = append(result, words[idx])
	}
	return strings.Join(result, " ")
}

func main() {
	original := "I want to invert the words of this string"
	result := InvertStringsWords(original)

	fmt.Println("Original:", original, "|", "Transformed:", result)
}
