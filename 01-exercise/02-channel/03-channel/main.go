package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 10)

	go func() {
		defer close(ch)

		// TODO: send all iterator values on channel without blocking

		for i := 0; i < 6; i++ {
			select {
			case ch <- i:
				fmt.Printf("Sending: %d\n", i)
			default:
				fmt.Println("Nothing to send")
				break
			}
		}
	}()

	for v := range ch {
		fmt.Printf("Received: %v\n", v)
	}

}
