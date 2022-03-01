package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go PrintEverySecond(&waitGroup, ctx)
	time.Sleep(5 * time.Second)
	cancel()
	waitGroup.Wait()
}

func PrintEverySecond(waitGroup *sync.WaitGroup, ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			waitGroup.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
