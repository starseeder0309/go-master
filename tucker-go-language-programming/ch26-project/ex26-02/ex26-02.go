package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	PrintFile("hamlet.txt")
}

func PrintFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("파일을 찾을 수 없습니다. ", fileName)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
