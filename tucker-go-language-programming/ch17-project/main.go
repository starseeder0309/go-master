package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var inputNumber int

	_, err := fmt.Scanln(&inputNumber)
	if err != nil {
		stdin.ReadString('\n')
	}

	return inputNumber, err
}

func main() {
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(100)

	count := 1
	for {
		fmt.Printf("숫자값을 입력하세요: ")

		inputNumber, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else {
			if inputNumber > randomNumber {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if inputNumber < randomNumber {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도한 횟수:", count)
				break
			}

			count++
		}
	}
}
