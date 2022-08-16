package main

import "fmt"

func Pow(number float64, power int) float64 {
	var result float64
	result = 1
	for idx := 0; idx < power; idx++ {
		result *= number
	}
	return result
}

func AsyncPow(out chan<- float64, number float64, power int) {
	out <- Pow(number, power)
}

func main() {
	out := make(chan float64)

	numbers := []float64{2, 4, 6, 8, 10}
	for _, number := range numbers {
		go AsyncPow(out, number, 2)
	}

	var sumNumbers float64
	sumNumbers = 0
	for idx := 0; idx < len(numbers); idx++ {
		sumNumbers += <-out
	}

	fmt.Println("Sum:", sumNumbers)
}
