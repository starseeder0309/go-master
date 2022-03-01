package main

import (
	"fmt"

	"tucker-go-language-programming/apxB-thinking-programming/exB-01/bankaccount"
)

func main() {
	account := bankaccount.NewAccount()
	account.Deposit(1000)
	fmt.Println(account.Balance())
}
