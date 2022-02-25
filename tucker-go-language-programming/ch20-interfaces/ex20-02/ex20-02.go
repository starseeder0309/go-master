package main

import "tucker-go-language-programming/ch20-interfaces/fedex"

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	// Fedex 전송 객체를 만든다.
	sender := &fedex.FedexSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}
