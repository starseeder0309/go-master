package main

import "fmt"

func main() {
	a := [5]int{32, 29, 78, 16, 81}

	for i, value := range a {
		fmt.Println(i, value)
	}

	fmt.Println()

	for _, value := range a {
		fmt.Println(value)
	}
}
