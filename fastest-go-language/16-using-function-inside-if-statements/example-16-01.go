package main

import (
	"fmt"
	"os"
)

func main() {
	var b []byte
	var err error

	b, err = os.ReadFile("./hello.txt")

	if err == nil {
		fmt.Printf("%s", b)
	}
}
