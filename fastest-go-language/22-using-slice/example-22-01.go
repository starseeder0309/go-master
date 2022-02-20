package main

import "fmt"

func main() {
	var a []int = make([]int, 5) // make 함수로 int 형에 길이가 5인 슬라이스에 공간을 할당한다.
	var b = make([]int, 5)       // 슬라이스를 선언할 때 자료형과 []를 생략할 수 있다.
	c := make([]int, 5)          // 슬라이스를 선언할 때 var 키워드, 자료형, []를 생략할 수 있다.

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
