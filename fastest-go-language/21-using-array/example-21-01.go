package main

import "fmt"

func main() {
	var a [5]int   // int형이며 길이가 5인 배열을 선언한다.
	a[2] = 7       // 배열의 세 번째 요소에 7을 대입한다.
	fmt.Println(a) // [0 0 7 0 0]
}
