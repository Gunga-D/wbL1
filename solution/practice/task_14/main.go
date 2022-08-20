package main

import "fmt"

func DetermineType(variable interface{}) string {
	switch t := variable.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	default:
		return fmt.Sprintf("unexpected type %T", t)
	}
}

func main() {
	var intVar int
	fmt.Println(DetermineType(intVar))

	var strVar string
	fmt.Println(DetermineType(strVar))

	var boolVar bool
	fmt.Println(DetermineType(boolVar))

	var floatVar float64
	fmt.Println(DetermineType(floatVar))
}
