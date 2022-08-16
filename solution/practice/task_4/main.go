package main

import (
	"context"
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

func AsyncShowRandomData(ctx context.Context, dataStream <-chan string) {
	for {
		select {
		case data := <-dataStream:
			fmt.Println(data)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	dataStream := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var workersNumber int
	fmt.Print("Workers number: ")
	fmt.Scanf("%d", &workersNumber)

	for workerIdx := 0; workerIdx < workersNumber; workerIdx++ {
		go AsyncShowRandomData(ctx, dataStream)
	}

	rand.Seed(time.Now().UnixNano())

	for {
		dataStream <- GenerateRandomString(4)
	}
}
