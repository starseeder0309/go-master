package main

import "fmt"

func main() {
	var r1 rune = '한'
	var r2 rune = '\ud55c'     // 한
	var r3 rune = '\U0000d55c' // 한

	fmt.Println(r1) // 54620
	fmt.Println(r2) // 54620
	fmt.Println(r3) // 54620
}
