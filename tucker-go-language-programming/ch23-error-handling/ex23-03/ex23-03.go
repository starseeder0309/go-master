package main

import "fmt"

type PasswordError struct {
	Length         int
	RequiredLength int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		return PasswordError{len(password), 8}
	}

	return nil
}

func main() {
	err := RegisterAccount("kimhyeonu", "0000")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf(
				"%v Length: %d, RequiredLength: %d\n",
				errInfo, errInfo.Length, errInfo.RequiredLength)
		}
	} else {
		fmt.Println("회원 가입에 성공했습니다.")
	}
}
