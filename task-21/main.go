package main

import "L1/task-21/adapters"

//21.Реализовать паттерн «адаптер» на любом примере.

func main() {
	adapter := adapters.NewAdapterJSONtoXMLSender()
	adapter.Send()
}
