package main

import (
	"context"
	"fmt"
	"time"
)

func closeGoroutineByChannelDefinition(stream <-chan struct{}) {
	for {
		_, ok := <-stream
		if !ok {
			return
		}

		fmt.Println("The goroutine will be closed by channel definition")
	}
}

func closeGoroutineByRange(stream <-chan struct{}) {
	for range stream {
		fmt.Println("The goroutine will be closed by range")
	}
}

func closeGoroutineByCancelChannel(cancel <-chan struct{}) {
	for {
		select {
		case <-cancel:
			return
		default:
			fmt.Println("The goroutine will be closed by cancel channel")
		}
	}
}

func closeGoroutineByContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("The goroutine will be closed by context")
		}
	}
}

func main() {
	streamForChanDef := make(chan struct{})
	go closeGoroutineByChannelDefinition(streamForChanDef)

	streamForRange := make(chan struct{})
	go closeGoroutineByRange(streamForRange)

	cancelChan := make(chan struct{})
	go closeGoroutineByCancelChannel(cancelChan)

	ctx, cancel := context.WithCancel(context.Background())
	go closeGoroutineByContext(ctx)

	time.Sleep(5 * time.Second)

	close(streamForChanDef)
	close(streamForRange)
	cancelChan <- struct{}{}
	cancel()

	for start := time.Now(); time.Since(start) < 2*time.Second; {
		fmt.Println("The main goroutine is working but other goroutines don't write anything...")
	}
}
