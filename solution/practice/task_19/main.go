package main

import "fmt"

func InvertString(input string) string {
	normInput := []rune(input)
	for left, right := 0, len(normInput)-1; left < right; left, right = left+1, right-1 {
		normInput[left], normInput[right] = normInput[right], normInput[left]
	}
	return string(normInput)
}

func main() {
	input := "I want to invert this string"
	result := InvertString(input)
	fmt.Println("Original string:", input, "|", "Inverted string:", result)
}
