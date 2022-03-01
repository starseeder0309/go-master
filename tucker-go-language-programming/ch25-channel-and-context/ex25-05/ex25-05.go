package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func square(ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Printf("Square: %d\n", n*n)
			time.Sleep(time.Second)
		case <-quit:
			fmt.Println("case: quit")
			waitGroup.Done()
			fmt.Println("after waitGroup.Done()")
			return
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan bool)

	waitGroup.Add(1)
	go square(ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	fmt.Println("loop outside")
	quit <- true

	waitGroup.Wait()
}
