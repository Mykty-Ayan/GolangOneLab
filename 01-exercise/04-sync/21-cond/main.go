package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
		cond.L.Lock()
		for len(sharedRsc) == 0 {
			cond.Wait()
		}
		cond.L.Unlock()

		fmt.Println(sharedRsc["rsc1"])
	}()

	// writes changes to sharedRsc

	sharedRsc["rsc1"] = "foo"
	wg.Wait()
}
