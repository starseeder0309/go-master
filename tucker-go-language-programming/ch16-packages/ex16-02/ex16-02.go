package main

import (
	"fmt"
	"tucker-go-language-programming/ch16-packages/ex16-02/publicpkg"
)

func main() {
	fmt.Println("PI:", publicpkg.PI)
	publicpkg.PublicFunc()

	var myInt publicpkg.MyInt = 10
	fmt.Println("myInt:", myInt)

	var myStruct = publicpkg.MyStruct{Age: 18}
	fmt.Println("myStruct:", myStruct)
}
