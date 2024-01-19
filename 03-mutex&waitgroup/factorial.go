package main

import (
	"fmt"
	"sync"
)

func main() {
	// f(10) = 55
	fmt.Println(f(10))
}

func f(n int) int {
	sum := 0
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}

	for i := 1; i < n+1; i++ {
		wg.Add(1)
		go func(j int) {
			// critical section
			mx.Lock()
			sum += j
			mx.Unlock()

			wg.Done()
		}(i)
	}
	wg.Wait()
	return sum
}
