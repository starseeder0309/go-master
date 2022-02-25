package main

import (
	"tucker-go-language-programming/ch20-interfaces/fedex"
	"tucker-go-language-programming/ch20-interfaces/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	fedexSender := &fedex.FedexSender{}
	SendBook("어린 왕자", fedexSender)
	SendBook("그리스인 조르바", fedexSender)

	koreaPostSender := &koreaPost.KoreaPostSender{}
	SendBook("어린 왕자", koreaPostSender)
	SendBook("그리스인 조르바", koreaPostSender)
}
