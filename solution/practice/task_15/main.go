package main

import (
	"fmt"
	"math/rand"
)

type Alph struct {
	letters []rune
}

func NewAlph() Alph {
	return Alph{letters: []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")}
}

func (a *Alph) GetRandomLetter() rune {
	return a.letters[rand.Intn(len(a.letters))]
}

func createHugeString(length int64) string {
	alph := NewAlph()

	var result []rune
	for idx := 0; idx < int(length); idx++ {
		result = append(result, alph.GetRandomLetter())
	}
	return string(result)
}

func someFunc() string {
	v := createHugeString(1 << 10)
	normV := []rune(v)

	justString := string(normV[:100])
	return justString
}

func main() {
	fmt.Println(someFunc())
}
