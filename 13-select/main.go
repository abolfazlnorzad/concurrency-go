package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second)
		ch1 <- 1
	}()

	go func() {
		ch2 <- 2
	}()

slct:
	for {
		select {
		case v1 := <-ch1:
			fmt.Println(v1)
		case v2 := <-ch2:
			fmt.Println(v2)
		case <-time.After(time.Second * 2):
			fmt.Println("finished.")
			break slct
		}
	}
}
