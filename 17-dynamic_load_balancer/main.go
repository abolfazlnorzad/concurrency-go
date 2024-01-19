package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	writer := func(ctx context.Context, wg *sync.WaitGroup) (<-chan int, <-chan struct{}) {
		dataCh := make(chan int)
		signalCh := make(chan struct{})
		go func() {
			defer close(dataCh)
			defer close(signalCh)
			defer wg.Done()
			ls := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40}
			for i := 0; i < len(ls); {
				var timeout time.Duration
				if i < 5 {
					timeout = time.Millisecond * 22
				} else if i < 10 {
					timeout = time.Millisecond * 44
				} else {
					timeout = time.Millisecond * 88
				}
				select {
				case dataCh <- ls[i]:
					i++
				case <-ctx.Done():
					return
				case <-time.After(timeout):
					signalCh <- struct{}{}
				}
			}
		}()
		return dataCh, signalCh
	}
	reader := func(ctx context.Context, dataCh <-chan int, wg *sync.WaitGroup, timeout time.Duration) {
		defer wg.Done()
		defer func() {
			fmt.Println("*******Bye.")
		}()
		for {
			time.Sleep(time.Millisecond * 77)
			select {
			case <-ctx.Done():
				return
			case i, ok := <-dataCh:
				if !ok {
					return
				}
				fmt.Println(i)
			case <-time.After(timeout):
				return

			}
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	dataCh, signalCh := writer(ctx, &wg)
	for range signalCh {
		wg.Add(1)
		fmt.Println("******New reader added")
		go reader(ctx, dataCh, &wg, time.Millisecond*100)
	}
	wg.Wait()
}
