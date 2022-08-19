package main

import "fmt"

func Push(to chan<- int, from []int) {
	for _, value := range from {
		to <- value
	}
	close(to)
}

func DoubleUp(from <-chan int, to chan<- int) {
	for value := range from {
		to <- value * value
	}
	close(to)
}

func Print(from <-chan int, trigger chan<- struct{}) {
	for value := range from {
		fmt.Println(value)
	}

	trigger <- struct{}{}
}

func main() {
	pushedStream := make(chan int)
	go Push(pushedStream, []int{
		1, 2, 3, 4, 5,
	})

	doubledStream := make(chan int)
	go DoubleUp(pushedStream, doubledStream)

	trigger := make(chan struct{})
	go Print(doubledStream, trigger)
	<-trigger
}
