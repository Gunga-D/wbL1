package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	total int64
}

func NewCounter() Counter {
	return Counter{total: 0}
}

func (c *Counter) Add(delta int64) {
	atomic.AddInt64(&c.total, delta)
}

func (c *Counter) Load() int64 {
	return atomic.LoadInt64(&c.total)
}

func main() {
	var wg sync.WaitGroup
	counter := NewCounter()

	for idx := 0; idx < 10; idx++ {
		wg.Add(1)
		go func() {
			counter.Add(1)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter.Load())
}
