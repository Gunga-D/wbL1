package main

import (
	"fmt"
	"time"
)

func sleep(delay time.Duration) {
	for start := time.Now(); time.Since(start) < delay; {
	}
}

func main() {
	fmt.Println("Current time:", time.Now())
	sleep(time.Second * 2)
	fmt.Println("After 2 seconds sleeping:", time.Now())
}
