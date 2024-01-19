package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	writer := func() {
		defer wg.Done()
		ls := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		for _, l := range ls {
			ch <- l
		}

		close(ch)
	}

	reader := func() {
		defer wg.Done()
		for i := range ch {
			fmt.Println("item : ", i)
			time.Sleep(time.Second)
		}
	}

	wg.Add(2)
	go writer()
	go reader()
	wg.Wait()

}
