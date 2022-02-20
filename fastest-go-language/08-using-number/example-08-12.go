package main

import (
	"fmt"
	"math"
)

func main() {
	var num1 uint8 = math.MaxUint8   // 255
	var num2 uint16 = math.MaxUint16 // 65535
	var num3 uint32 = math.MaxUint32 // 4294967295
	var num4 uint64 = math.MaxUint64 // 18446744073709551615

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(num4)
}
