package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	writer := func() {
		ls := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

		for _, l := range ls {
			ch <- l
		}
		close(ch)
		wg.Done()
	}

	reader := func() {
		/***
		@see deadlock!

		We need to loop in buffered channels, so every time we write data to the channel, we also read from it once.
		Here, we are writing multiple times but only reading once. We need to loop.
		i := <-ch
		fmt.Println(i)
		wg.Done()
		*/

		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()

	}

	wg.Add(1)
	go writer()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go reader()
	}
	wg.Wait()
}
