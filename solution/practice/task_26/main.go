package main

import (
	"fmt"
	"strings"
)

func IsStringUnique(input string) bool {
	buffer := make(map[rune]struct{})

	normString := strings.ToLower(input)
	for _, symb := range normString {
		if _, found := buffer[symb]; found {
			return false
		}

		buffer[symb] = struct{}{}
	}
	return true
}

func main() {
	input := "This string is not unique"
	fmt.Println("Checking string:", input, "| Result:", IsStringUnique(input))

	input = "Thisrngouqe"
	fmt.Println("Checking string:", input, "| Result:", IsStringUnique(input))
}
