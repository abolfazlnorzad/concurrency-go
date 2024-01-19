package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		c := 0
		go func() {
			fmt.Printf("c is :%d \n", c)
		}()

		go func() {
			c++
		}()
		time.Sleep(10 * time.Millisecond)
	}

}
