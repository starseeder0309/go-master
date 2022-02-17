package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println()

	for i := 5; i > 0; i-- {
		fmt.Println(i)
	}
}
