package main

import "fmt"

func main() {
	var a int = 1

	if a == 1 {
		goto ERROR1
	}

	if a == 2 {
		goto ERROR1
	}

	if a == 3 {
		goto ERROR2
	}

	fmt.Println(a)
	fmt.Println("Success")

	return

ERROR1: // 에러 처리 1
	fmt.Println("Error 1")
	return

ERROR2: // 에러 처리 2
	fmt.Println("Error 2")
	return
}
