// generator() -> square() -> print

package main

import (
	"fmt"
	"sync"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	// Implement fan-in
	// merge a list of channels to a single channel
	var wg sync.WaitGroup
	out := make(chan int)

	for _, channel := range cs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()

			for num := range ch {
				out <- num
			}
		}(channel)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := generator(2, 3)

	// TODO: fan out square stage to run two instances.
	squareChannel1 := square(in)
	squareChannel2 := square(in)
	// TODO: fan in the results of square stages.
	for num := range merge(squareChannel1, squareChannel2) {
		fmt.Println("Received from fan in ", num)
	}
}
