package main

import "fmt"

func binaryFind(from []int, what int) int {
	left := 0
	right := len(from) - 1

	for left <= right {
		middle := left + (right-left)/2

		if from[middle] == what {
			return middle
		}

		if from[middle] < what {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}

func main() {
	array := []int{1, 2, 3, 4, 5}

	fmt.Println("Array:", array, "|", "Index of 3:", binaryFind(array, 3))
}
