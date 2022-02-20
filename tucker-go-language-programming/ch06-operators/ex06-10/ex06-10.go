package main

import "fmt"

func main() {
	var a int = 10
	var b int = 20

	a, b = b, a // a와 b 값을 서로 바꾼다.

	fmt.Println(a, b)
}
