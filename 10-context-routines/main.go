package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	result := make(chan string, 2)

	wg.Add(2)

	go worker(ctx, 1, result, &wg)
	go worker(ctx, 2, result, &wg)

	go func() {
		wg.Wait()
		close(result)
	}()

	go func() {
		time.Sleep(350 * time.Millisecond)
		cancel()
	}()

	for message := range result {
		fmt.Println(message)
	}
}

func worker(ctx context.Context, id int, result chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			subString := fmt.Sprintf("worker %d stopped", id)
			result <- subString
			return

		case <-time.After(100 * time.Millisecond):
			subString := fmt.Sprintf("worker %d working", id)
			result <- subString
		}
	}
}
