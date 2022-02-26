package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	line, _ := reader.ReadString('\n')

	return line, nil
}

func WriteFile(fileName string, line string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = fmt.Fprintln(file, line)

	return err
}

const fileName string = "data.txt"

func main() {
	line, err := ReadFile(fileName)

	if err != nil {
		err = WriteFile(fileName, "This is WriteFile.")
		if err != nil {
			fmt.Println("파일 생성에 실패했습니다.", err)
			return
		}

		line, err = ReadFile(fileName)
		if err != nil {
			fmt.Println("파일 읽기에 실패했습니다.", err)
			return
		}
	}

	fmt.Println("파일 내용:", line)
}
