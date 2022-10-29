package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func(a, b int) {
		defer close(ch)
		ch <- a + b
	}(1, 2)

	sum := <-ch
	// TODO: get the value computed from goroutine
	fmt.Printf("computed value %v\n", sum)
}
