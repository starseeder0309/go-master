package main

import "fmt"

func main() {
	var a int16 = 3456
	var c int8 = int8(a) // int16 타입에서 int8 타입으로 변환

	fmt.Println(a)
	fmt.Println(c) // int8 타입인 c 값 출력
}
