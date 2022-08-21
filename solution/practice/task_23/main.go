package main

import "fmt"

func DeleteElementFrom(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func main() {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := DeleteElementFrom(input, 3)

	fmt.Println("Original:", input, "|", "After deleting element with index 3", result)
}
