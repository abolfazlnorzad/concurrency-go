package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	i := 0
	j := 0
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i == 0 {
			time.Sleep(time.Second)
		}
		j = 1
		fmt.Println("j=1 , done!")
	}()

	go func() {
		defer wg.Done()
		for j == 0 {
			time.Sleep(time.Second)
		}
		i = 1
		fmt.Println("i=1 , done!")
	}()
	wg.Wait()
}
