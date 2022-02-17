package main

import "fmt"

// <- (채널 수신 연산)
// 채널에 값을 보내거나 값을 가져온다.
func main() {
	c := make(chan int)

	go func() {
		c <- 1 // 채널 c에 1을 보낸다.
	}()

	a := <-c // 채널 c에서 값을 가져와서 a에 대입

	fmt.Println(a)
}
