package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3, 10)
	slice3 := make([]int, 10)

	count1 := copy(slice2, slice1)
	count2 := copy(slice3, slice1)

	fmt.Println(count1, slice2)
	fmt.Println(count2, slice3)
}
