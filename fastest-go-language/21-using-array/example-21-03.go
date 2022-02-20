package main

import "fmt"

func main() {
	a := [5]int{12, 29, 78, 16, 81}
	b := [5]int{
		32,
		29,
		78,
		16,
		81, // 여러 줄로 나열할 때는 마지막 요소에 콤마(,)를 붙인다.
	}

	c := [...]int{32, 29, 78, 16, 81}
	d := [...]string{"Maria", "Andrew", "John"}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
