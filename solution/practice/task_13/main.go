package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 1
	b := 2
	fmt.Println("A:", a, "|", "B:", b)

	swap(&a, &b)
	fmt.Println("A:", a, "|", "B:", b)
}
