package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin) // 표준 입력을 읽는 객체

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)

	if err != nil { // 에러 발생 시
		fmt.Println(err)       // 에러 출력
		stdin.ReadString('\n') // 표준 입력 스트림 지우기
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(n, a, b)
	}
}
