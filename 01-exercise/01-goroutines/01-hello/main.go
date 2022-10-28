package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func fun(s string) {

	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(1 * time.Millisecond)
	}
}

func call(f func(s string), s string) {
	defer wg.Done()
	f(s)
}

func main() {
	// Direct call
	fun("direct call")

	// TODO: write goroutine with different variants for function call.

	// goroutine function call
	wg.Add(1)
	go fun("function call")
	wg.Done()

	// goroutine with anonymous function
	wg.Add(1)
	go func() {
		defer wg.Done()
		fun("anonymous function")
	}()

	// goroutine with function value call
	wg.Add(1)
	go call(fun, "function value call")

	// wait for goroutines to end
	wg.Wait()

	fmt.Println("done..")
}
