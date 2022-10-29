package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		defer close(ch)
		time.Sleep(2 * time.Second)
		ch <- "one"
	}()

	// TODO: implement timeout for recv on channel ch
	select {
	case s := <-ch:
		fmt.Println("Received " + s + " from first channel")
	case <-time.After(time.Second * 5):
		fmt.Println("Timeout")
	}

	m := <-ch
	fmt.Println(m)
}
