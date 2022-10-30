package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {

	// TODO: generator -  generates integers in a separate goroutine and
	// sends them to the returned channel.
	// The callers of gen need to cancel the goroutine once
	// they consume 5th integer value
	// so that internal goroutine
	// started by gen is not leaked.
	var wg sync.WaitGroup
	wg.Add(1)
	generator := func(ctx context.Context) <-chan int {
		out := make(chan int)
		defer wg.Done()
		go func() {
			defer close(out)
			for i := 0; i < 10000; i += 1 {
				select {
				case out <- i:
				case <-ctx.Done():
					return
				}
			}
		}()

		return out
	}

	// Create a context that is cancellable.
	ctx, cancel := context.WithCancel(context.Background())

	for num := range generator(ctx) {
		fmt.Println(num)
		if num == 4 {
			cancel()
		}

	}
	wg.Wait()
}
