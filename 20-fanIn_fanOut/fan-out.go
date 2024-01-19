package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)

	tasks := func() {
		ls := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
		for _, s := range ls {
			ch <- s
		}

		close(ch)
		wg.Done()
	}

	workers := func() {
		for v := range ch {
			fmt.Println(v)
		}
		wg.Done()
	}

	go tasks()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go workers()
	}

	wg.Wait()
}
