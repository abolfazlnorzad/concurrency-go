package main

import "fmt"

func main() {
	readOnlyCh := Generator()
	for i := range readOnlyCh {
		fmt.Println(i)
	}
}

func Generator() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		ls := []int{1, 2, 3, 4, 5, 6, 7}
		for _, l := range ls {
			ch <- l
		}
	}()
	return ch
}
