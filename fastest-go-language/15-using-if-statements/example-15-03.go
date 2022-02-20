package main

import "fmt"

func main() {
	i := 6

	if i >= 10 {
		fmt.Println("i는 10 이상의 값입니다.")
	} else if i >= 5 && i < 10 {
		fmt.Println("i는 5 이상, 10 미만의 값입니다.")
	} else {
		fmt.Println("i는 0 이상, 5 미만의 값입니다.")
	}
}
