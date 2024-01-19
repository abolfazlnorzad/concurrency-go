package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		ls := []int{1, 2, 3, 4, 5}
		for _, l := range ls {
			ch <- l
		}
	}()

	go func() {
		defer wg.Done()
		ls := []int{6, 7, 8, 9, 10}
		for _, l := range ls {
			ch <- l
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}

}
