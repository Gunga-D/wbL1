package main

import "fmt"

func main() {
	temperatures := []float64{
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5,
	}
	groups := map[int][]float64{}

	for _, temp := range temperatures {
		groupingTemp := int(temp/10) * 10
		groups[groupingTemp] = append(groups[groupingTemp], temp)
	}
	fmt.Println("Groups:", groups)
}
