package main

import (
	"fmt"
)

func Pow(number float64, power int) float64 {
	var result float64
	result = 1
	for idx := 0; idx < power; idx++ {
		result *= number
	}
	return result
}

func ShowPoweredNumber(number float64, power int) {
	fmt.Printf("%f -> %f\n", number, Pow(number, power))
}

func AsyncShowPoweredNumber(done chan<- struct{}, number float64, power int) {
	ShowPoweredNumber(number, power)
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})

	numbers := []float64{2, 4, 6, 8, 10}
	for _, number := range numbers {
		go AsyncShowPoweredNumber(done, number, 2)
	}

	for idx := 0; idx < len(numbers); idx++ {
		<-done
	}
}
