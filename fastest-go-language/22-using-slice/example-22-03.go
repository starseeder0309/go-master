package main

import "fmt"

func main() {
	a := make([]int, 5, 10)

	fmt.Println(len(a))
	fmt.Println(cap(a))

	// 용량이 길이보다 크더라도 길이를 벗어난 인덱스에는 접근할 수 없다.
	fmt.Println(a[4]) // 0: make 함수를 사용하면 슬라이스의 요소는 모두 0으로 초기화된다.
	fmt.Println(a[5]) // 길이를 벗어난 인덱스에 접근했으므로 런타임 에러가 발생한다.
	fmt.Println(a[5]) // 길이를 벗어난 인덱스에 접근했으므로 런타임 에러가 발생한다.
}
