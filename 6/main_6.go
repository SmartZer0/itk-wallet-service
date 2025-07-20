package main

import (
	"fmt"
	"math/rand"
	"time"
)

func startRandomGenerator() <-chan int {
	ch := make(chan int)
	go func() {
		for {
			ch <- rand.Intn(1000)
		}
	}()
	return ch
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := startRandomGenerator()

	for i := 0; i < 5; i++ {
		fmt.Println("Получено число:", <-ch)
	}
}
