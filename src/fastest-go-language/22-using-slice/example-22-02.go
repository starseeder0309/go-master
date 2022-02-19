package main

import "fmt"

func main() {
	a := []int{32, 29, 78, 16, 81} // 슬라이스를 생성하면서 초기화
	b := []int{
		32,
		29,
		78,
		16,
		81,
	}

	fmt.Println(a)
	fmt.Println(b)
}
