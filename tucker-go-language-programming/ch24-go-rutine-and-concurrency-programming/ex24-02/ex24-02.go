package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}

	fmt.Printf("%d부터 %d까지의 합계는 %d입니다.\n", a, b, sum)

	waitGroup.Done()
}

func main() {
	waitGroup.Add(10)

	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000000)
	}

	waitGroup.Wait()

	fmt.Println("모든 계산이 완료됐습니다.")
}
