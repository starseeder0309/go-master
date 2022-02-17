package main

import "fmt"

func main() {
	const (
		a1 = 1 << iota // a1 == 1 (1 << 0)
		b1 = 1 << iota // b1 == 2 (1 << 1)
		c1 = 1 << iota // c1 == 4 (1 << 2)
		d1 = 1 << iota // d1 == 8 (1 << 3)
	)

	const (
		a2 = iota * 30 // a2 == 0 (0 * 30)
		b2 = iota * 30 // b2 == 30 (1 * 30)
		c2 = iota * 30 // c2 == 60 (2 * 30)
		d2 = iota * 30 // d2 == 90 (3 * 30)
	)

	fmt.Println(a1, b1, c1, d1)
	fmt.Println(a2, b2, c2, d2)
}
