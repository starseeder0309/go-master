package main

import (
	"fmt"
	"os"
)

func main() {
	if b, err := os.ReadFile("./hello.txt"); err == nil {
		fmt.Printf("%s", b)
	}
}
