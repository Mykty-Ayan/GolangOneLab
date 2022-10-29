package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(cnl chan<- int) {
		// TODO: send iterator over channel
		for i := 0; i < 6; i++ {
			ch <- i
		}
		close(cnl)
	}(ch)

	// TODO: range over channel to recv values
	for num := range ch {
		fmt.Printf("Received %d from channel\n", num)
	}

}
