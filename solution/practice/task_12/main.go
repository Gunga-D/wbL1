package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})
	for _, val := range arr {
		set[val] = struct{}{}
	}
	fmt.Println("Set:", set)
}
