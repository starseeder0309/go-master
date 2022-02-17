package main

import "fmt"

// & (참조 연산)
// 현재 변수의 메모리 주소를 구한다.
func main() {
	a := 1
	b := &a        // a의 메모리 주소를 b에 대입
	fmt.Println(b) // 1: a에 저장된 메모리에 접근하여 값을 가져온다.
}
