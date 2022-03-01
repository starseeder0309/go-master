package main

import "fmt"

func fibonacci1(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {
		return n
	}

	return fibonacci1(n-1) + fibonacci1(n-2)
}

func fibonacci2(n int) int {
	if n < 0 {
		return 0
	}

	if n < 2 {
		return n
	}

	first := 1
	second := 0
	result := 0
	for i := 2; i <= n; i++ {
		result = first + second
		second = first
		first = result
	}

	return result
}

func main() {
	fmt.Println(fibonacci1(13))
	fmt.Println(fibonacci2(13))
}
