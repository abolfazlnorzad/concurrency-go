package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	files := make(chan int)
	dirs := make(chan string)
	tokens := make(chan int, 10) // number of goroutines.
	wg := &sync.WaitGroup{}

	countFiles := func(path string) {
		defer func() {
			wg.Done()
			<-tokens
		}()
		entities, err := os.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, entity := range entities {
			if entity.IsDir() {
				dirs <- fmt.Sprintf("%s/%s", path, entity.Name())
			} else {
				files <- 1
			}

		}
	}

	go func() {
		for dir := range dirs {
			wg.Add(1)
			tokens <- 1
			go func(d string) {
				countFiles(d)
			}(dir)
		}
	}()
	root := "."
	dirs <- root

	go func() {
		wg.Wait()
		close(files)
		close(dirs)
	}()

	cnt := 0

	for file := range files {
		cnt += file
	}

	fmt.Println(cnt)

}
