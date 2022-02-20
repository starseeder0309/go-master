package main

import "fmt"

const Y int = 3

func main() {
	// x := 5
	// a := [x]int{1, 2, 3, 4, 5} // 에러 발생 - 변수는 배열 길이로 사용할 수 없음
	b := [Y]int{1, 2, 3}
	var c [10]int

	fmt.Println("b:", b)
	fmt.Println("c:", c)
}
