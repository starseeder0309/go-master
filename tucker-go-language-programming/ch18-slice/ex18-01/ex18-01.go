package main

import "fmt"

func main() {
	var slice []int

	if len(slice) == 0 { // 슬라이스의 길이가 0인지 확인한다.
		fmt.Println("slice is empty.", slice)
	}

	slice[1] = 10 // panic: runtime error: index out of range [1] with length 0
	fmt.Println(slice)
}
