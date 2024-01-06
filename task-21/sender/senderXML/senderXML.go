package senderxml

//Отправка данных в БД работающая только с XML форматом

import (
	"L1/task-21/sender"
	"fmt"
)

type SenderXML struct {
}

func NewSenderXML() sender.Sender {
	return &SenderXML{}
}

func (s *SenderXML) Send() {
	fmt.Println("Send XML to XML_DB")
}
