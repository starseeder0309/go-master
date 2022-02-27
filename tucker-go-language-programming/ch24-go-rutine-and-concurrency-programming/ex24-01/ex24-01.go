package main

import (
	"fmt"
	"time"
)

func PrintKorean() {
	letters := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, letter := range letters {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", letter)
	}
}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go PrintKorean()
	go PrintNumbers()

	time.Sleep(3 * time.Second)
}
