package main

import (
	"fmt"
	"time"
)

/*
**
@see

In places where we have fast writers and slow readers, using buffered channels may not be advisable.
This is because, in such cases, we often risk reaching high channel capacities, leading to a waste of system resources.
Buffered channels are typically employed when we know exactly how many elements a writer is going to write into the channel,
or to limit the number of goroutines.
*/
func main() {
	ch := make(chan int, 5)

	writer := func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		fmt.Println("writing 5 element.")
		ch <- 6
		fmt.Println("writing 6 element.")

		close(ch)
	}
	reader := func() {
		time.Sleep(time.Second)
		for i := range ch {
			fmt.Println(i)
		}
	}
	go writer()
	go reader()
	time.Sleep(2 * time.Second)
}
