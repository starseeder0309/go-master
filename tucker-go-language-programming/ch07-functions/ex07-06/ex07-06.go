package main

import "fmt"

func PrintNo(n int) {
	if n == 0 {
		return
	}

	fmt.Println(n)
	PrintNo(n - 1)
	fmt.Println("After", n)
}

func main() {
	PrintNo(3)
}
