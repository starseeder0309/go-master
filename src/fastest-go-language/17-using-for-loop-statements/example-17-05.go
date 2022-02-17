package main

import "fmt"

func main() {
Loop: // Loop 레이블 지정. 레이블과 for 키워드 사이에 다른 코드가 있으면 안 됨.
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break Loop
			}

			fmt.Println(i, j)
		}
	}

	fmt.Println("Hello, world!")
}
