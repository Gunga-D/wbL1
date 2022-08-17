package main

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenerateRandomString(stringLength int) string {
	result := make([]rune, stringLength)
	for idx := range result {
		result[idx] = letters[rand.Intn(stringLength)]
	}
	return string(result)
}

func AsyncSender(dataStream chan<- string, duration int) {
	for start := time.Now(); time.Since(start) < time.Second*time.Duration(duration); {
		dataStream <- GenerateRandomString(3)
	}
	close(dataStream)
}

func main() {
	dataStream := make(chan string)

	go AsyncSender(dataStream, 3)

	for data := range dataStream {
		fmt.Println(data)
	}
}
