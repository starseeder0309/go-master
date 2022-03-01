package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func square(ch chan int) {
	for n := range ch {
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}

	waitGroup.Done()
}

func main() {
	ch := make(chan int)

	waitGroup.Add(1)
	go square(ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	close(ch)
	waitGroup.Wait()
}
