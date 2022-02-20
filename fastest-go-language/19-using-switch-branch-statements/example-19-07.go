package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	switch i := rand.Intn(10); {
	case i >= 3 && i < 6:
		fmt.Print("3 이상, 6 미만의 값")
	case i == 9:
		fmt.Println("9")
	default:
		fmt.Println(i)
	}
}
