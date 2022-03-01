package main

import (
	"context"
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(1)
	ctx := context.WithValue(context.Background(), "number", 9)
	go square(ctx)
	waitGroup.Wait()
}

func square(ctx context.Context) {
	if value := ctx.Value("number"); value != nil {
		number := value.(int)
		fmt.Printf("Square: %d", number*number)
	}

	waitGroup.Done()
}
