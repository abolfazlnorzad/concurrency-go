package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	sharedValue := 1
	mux := sync.RWMutex{}
	wg := sync.WaitGroup{}
	read := func() {
		defer wg.Done()
		// wg.Add(1) // it's bug.
		start := time.Now()
		mux.RLock()
		time.Sleep(time.Millisecond * 100)
		mux.RUnlock()
		fmt.Printf("its took me %d milisecond to do my job(read). \n", time.Now().Sub(start).Milliseconds())

	}
	write := func() {
		defer wg.Done()
		// wg.Add(1) // it's bug.
		start := time.Now()
		for mux.TryLock() == false {
			time.Sleep(time.Millisecond * 10)
		}
		sharedValue++
		mux.Unlock()
		fmt.Printf("its took me %d milisecond to do my job(write). \n", time.Now().Sub(start).Milliseconds())
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
}
