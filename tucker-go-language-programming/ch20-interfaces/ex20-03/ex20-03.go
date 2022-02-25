package main

import (
	"tucker-go-language-programming/ch20-interfaces/fedex"
	"tucker-go-language-programming/ch20-interfaces/koreaPost"
)

func SendBook(name string, sender *fedex.FedexSender) {
	sender.Send(name)
}

func main() {
	// 우체국 전송 객체를 만든다.
	sender := &koreaPost.KoreaPostSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
}
