package main

import (
	"fmt"
	"sync"
	"time"
)

func square(waitGroup *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated!")
			waitGroup.Done()
			return
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var waitGroup sync.WaitGroup
	ch := make(chan int)

	waitGroup.Add(1)
	go square(&waitGroup, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	waitGroup.Wait()
}
