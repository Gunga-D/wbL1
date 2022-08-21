package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func partition(array []int, low int, high int) int {
	pivot := array[high]

	i := low - 1
	for j := low; j < high; j++ {
		if array[j] <= pivot {
			i++
			swap(&array[i], &array[j])
		}
	}

	swap(&array[i+1], &array[high])

	return i + 1
}

func quicksort(array []int, low int, high int) {
	if low < high {
		pivot := partition(array, low, high)

		quicksort(array, low, pivot-1)
		quicksort(array, pivot+1, high)
	}
}

func main() {
	array := []int{5, 4, 3, 2, 1}
	quicksort(array, 0, len(array)-1)
	fmt.Println("Sorted array: ", array)
}
