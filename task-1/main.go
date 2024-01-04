package main

import "fmt"

//1. Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

// Создание структуры Human
type Human struct {
	name string
	age  uint
}

func (h *Human) Introduce() {
	fmt.Printf("Hi, my name is %s. I am %d old.\n", h.name, h.age)
}

// Создание структуры Action, куда добавляем анонимное поле типа Human
type ActionAnon struct {
	Human
}

// Создание структуры Action, куда добавляем поле типа Human
type Action struct {
	human Human
}

func main() {
	actionAnon := ActionAnon{Human{"Andrew", 26}}
	actionAnon.Introduce()

	action := Action{human: Human{"Liza", 21}}
	action.human.Introduce()
}
