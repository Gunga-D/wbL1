package main

import "fmt"

func main() {
	firstSet := map[int]struct{}{
		2:  struct{}{},
		3:  struct{}{},
		8:  struct{}{},
		21: struct{}{},
	}
	secondSet := map[int]struct{}{
		8: struct{}{},
		6: struct{}{},
		3: struct{}{},
		1: struct{}{},
	}

	intesection := []int{}
	for value, _ := range firstSet {
		_, found := secondSet[value]
		if found {
			intesection = append(intesection, value)
		}
	}
	fmt.Println("Intersection:", intesection)
}
