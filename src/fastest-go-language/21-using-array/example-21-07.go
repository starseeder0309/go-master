package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	b := a

	// 배열 변수는 배열 전체를 뜻하며,
	// 첫 번째 요소의 포인터가 아니다.
	// 배열을 다른 변수에 대입하면 배열 전체가 복사된다.
	fmt.Println(a) // [1, 2, 3, 4, 5]
	fmt.Println(b) // [1, 2, 3, 4, 5]
}
