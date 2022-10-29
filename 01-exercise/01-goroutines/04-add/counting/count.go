package counting

import (
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateNumbers - random number generation
func GenerateNumbers(max int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, max)
	for i := 0; i < max; i++ {
		numbers[i] = rand.Intn(10)
	}
	return numbers
}

// Add - sequential code to add numbers
func Add(numbers []int) int64 {
	var sum int64
	for _, n := range numbers {
		sum += int64(n)
	}
	return sum
}

//TODO: complete the concurrent version of add function.

// AddConcurrent - concurrent code to add numbers
func AddConcurrent(numbers []int) int64 {
	var sum int64
	var wg sync.WaitGroup

	// Utilize all cores on machine
	numCPU := runtime.NumCPU()
	stride := len(numbers) / numCPU

	// Divide the input into parts
	wg.Add(numCPU)
	// Run computation for each part in seperate goroutine.
	for i := 0; i < numCPU; i += 1 {
		go func(cpu int) {
			defer wg.Done()
			start := cpu * stride
			end := start + stride
			if cpu == numCPU-1 {
				end = len(numbers)
			}
			var partialSum int
			for _, n := range numbers[start:end] {
				partialSum += n
			}

			atomic.AddInt64(&sum, int64(partialSum))

		}(i)
	}

	// Add part sum to cummulative sum
	wg.Wait()
	return sum
}
