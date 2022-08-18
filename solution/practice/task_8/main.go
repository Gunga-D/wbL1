package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input int64
	fmt.Print("Input number: ")
	fmt.Scan(&input)

	inputBitRepr := strconv.FormatInt(input, 2)

	var i int
	fmt.Print("Index of changing bit: ")
	fmt.Scan(&i)

	if inputBitRepr[i] == '1' {
		inputBitRepr = inputBitRepr[0:i] + "0" + inputBitRepr[i+1:]
	} else {
		inputBitRepr = inputBitRepr[0:i] + "1" + inputBitRepr[i+1:]
	}

	result, err := strconv.ParseInt(inputBitRepr, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
