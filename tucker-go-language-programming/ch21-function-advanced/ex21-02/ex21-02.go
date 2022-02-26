package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file.")
		return
	}

	defer fmt.Println("반드시 호출됩니다.")
	defer file.Close()
	defer fmt.Println("파일을 닫습니다.")

	fmt.Println("파일에 Hello World를 씁니다.")
	fmt.Fprintln(file, "Hello World")
}
