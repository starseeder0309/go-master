package main

import "fmt"

func main() {
	var a [5]int = [5]int{32, 29, 78, 16, 81} // int 형이며 길이가 5인 배열을 선언하고 초기화한다.
	var b = [5]int{32, 29, 78, 16, 81}        // 배열을 선언할 때 자료형과 길이를 생략한다.
	c := [5]int{32, 29, 78, 16, 81}           // 배열을 선언할 때 var 키워드, 자료형과 같이 생략한다.

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
