package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func AsyncWriteToMap(waiter *sync.WaitGroup, locker *sync.Mutex, hashMap map[int]int, writtingCount int) {
	for idx := 0; idx < writtingCount; idx++ {
		randomKey := rand.Intn(100)
		randomValue := rand.Intn(100)

		locker.Lock()
		hashMap[randomKey] = randomValue
		locker.Unlock()
	}
	waiter.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	hashMap := make(map[int]int)

	var waiter sync.WaitGroup
	waiter.Add(3)

	var locker sync.Mutex

	go AsyncWriteToMap(&waiter, &locker, hashMap, 10)
	go AsyncWriteToMap(&waiter, &locker, hashMap, 10)
	go AsyncWriteToMap(&waiter, &locker, hashMap, 10)

	waiter.Wait()

	fieldNumber := 1
	for key, value := range hashMap {
		fmt.Println("Number:", fieldNumber, "|", "Key:", key, "|", "Value:", value)
		fieldNumber++
	}
}
