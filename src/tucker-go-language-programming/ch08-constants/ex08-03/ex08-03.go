package main

import "fmt"

// 상숫값에 코드를 부여한다.
const Pig int = 0
const Cow int = 1
const Chicken int = 2

// 코드 값에 따라서 다른 텍스트를 출력한다.
func PrintAnimal(animal int) {
	if animal == Pig {
		fmt.Println("꿀꿀")
	} else if animal == Cow {
		fmt.Println("음메")
	} else if animal == Chicken {
		fmt.Println("꼬끼오")
	} else {
		fmt.Println("...")
	}
}

func main() {
	PrintAnimal(Pig)
	PrintAnimal(Cow)
	PrintAnimal(Chicken)
}
