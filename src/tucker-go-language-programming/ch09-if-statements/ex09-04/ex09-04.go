package main

import "fmt"

var count int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", count)
	count++
	return count
}

func main() {
	if false && IncreaseAndReturn() < 5 {
		fmt.Println("1 증가")
	}

	if true || IncreaseAndReturn() < 5 {
		fmt.Print("2 증가")
	}

	fmt.Println("count : ", count)
}
