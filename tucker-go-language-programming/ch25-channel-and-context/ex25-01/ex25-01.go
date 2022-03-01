package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func square(ch chan int) {
	fmt.Println("called square()")
	number := <-ch
	time.Sleep(time.Second)
	fmt.Printf("Square: %d\n", number*number)
	waitGroup.Done()
}

func main() {
	ch := make(chan int)

	waitGroup.Add(1)
	go square(ch)
	ch <- 9

	waitGroup.Wait()
}
