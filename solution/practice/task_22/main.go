package main

import (
	"fmt"
	"math/big"
)

func Add(firstNumber *big.Int, secondNumber *big.Int) *big.Int {
	return big.NewInt(0).Add(firstNumber, secondNumber)
}

func Sub(firstNumber *big.Int, secondNumber *big.Int) *big.Int {
	return big.NewInt(0).Sub(firstNumber, secondNumber)
}

func Multiply(firstNumber *big.Int, secondNumber *big.Int) *big.Int {
	return big.NewInt(0).Mul(firstNumber, secondNumber)
}

func Divide(firstNumber *big.Int, secondNumber *big.Int) *big.Int {
	return big.NewInt(0).Div(firstNumber, secondNumber)
}

func main() {
	firstNumber := big.NewInt(1 << 22)
	secondNumber := big.NewInt(1 << 21)

	fmt.Println("Result of", firstNumber, "+", secondNumber, "=", Add(firstNumber, secondNumber))
	fmt.Println("Result of", firstNumber, "-", secondNumber, "=", Sub(firstNumber, secondNumber))
	fmt.Println("Result of", firstNumber, "*", secondNumber, "=", Multiply(firstNumber, secondNumber))
	fmt.Println("Result of", firstNumber, "/", secondNumber, "=", Divide(firstNumber, secondNumber))
}
