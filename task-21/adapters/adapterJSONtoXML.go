package adapters

//Адаптер для условной конфертации JSON файла в XML для последующей передачи в senderxml и отправки в БД

import (
	"L1/task-21/sender"
	senderxml "L1/task-21/sender/senderXML"
	"fmt"
)

type AdapterJSONtoXMLSender struct {
	xmlS sender.Sender
}

func NewAdapterJSONtoXMLSender() AdapterJSONtoXMLSender {
	return AdapterJSONtoXMLSender{xmlS: senderxml.NewSenderXML()}
}

func (n *AdapterJSONtoXMLSender) Send() {
	fmt.Println("Converting JSON to XML...")
	n.xmlS.Send()
}
