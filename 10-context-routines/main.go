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

	wg.Add(1)

	go worker(ctx, &wg)

	time.Sleep(350 * time.Millisecond)
	cancel()

	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker stopped")
			return

		case <-time.After(100 * time.Millisecond):
			fmt.Println("worker working")
		}
	}
}
