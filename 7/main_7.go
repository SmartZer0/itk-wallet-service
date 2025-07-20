package main

import (
	"fmt"
	"sync"
)

func merge(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(channels))
	for _, ch := range channels {
		go func(c <-chan int) {
			defer wg.Done()
			for val := range c {
				out <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// Пример использования
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()

	go func() {
		defer close(ch2)
		ch2 <- 3
		ch2 <- 4
	}()

	for val := range merge(ch1, ch2) {
		fmt.Println(val)
	}
}
