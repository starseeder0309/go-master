package main

import "fmt"

func main() {
	var a int = 1

	if a == 1 {
		fmt.Println("Error 1") // 에러 1 상황
		return
	}

	if a == 2 {
		fmt.Println("Error 2") // 에러 2 상황
	}

	if a == 3 {
		fmt.Println("Error 1") // 에러 1 상황. 중복 코드
		return
	}

	fmt.Println(a)
	fmt.Println("Success") // 정상 실행

	return
}
