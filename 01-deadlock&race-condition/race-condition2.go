package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := 0
	wg := sync.WaitGroup{}
	wg.Add(2)
	longRunning := func() {
		if c == 0 {
			defer func() { c = 0 }()
			c = 1
			fmt.Println("i'm greedy.")
			time.Sleep(10 * time.Nanosecond)
		}
	}

	shortRunning := func() {
		if c == 0 {
			defer func() { c = 0 }()
			c = 1
			fmt.Println("i'm polite")
		} else {
			time.Sleep(time.Millisecond)
		}
	}

	go func() {
		for {
			longRunning()
		}
	}()

	go func() {
		time.Sleep(time.Millisecond)
		for {
			shortRunning()
		}
	}()
	wg.Wait()
}
