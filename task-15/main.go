package main

//15.К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?

//Проблема вышеуказанного кода состоит в том что выделяется слишком большой объем памяти для строки.
//Где в строке 11 будет использована всего лишь 10-ая часть от выделенной памяти

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

//Приведите корректный пример реализации.

//Если в 11 строке нам так и будет нужно создавать массив из 100 ячеек то достаточно исправить строку 10
//Но лучше в функцию  func someFunc() передавать параметры на создание большой строки и обычной строки))

var justString string

func someFunc() {
	v := createHugeString(1 << 7)
	justString = v[:100]
}

func main() {
	someFunc()
}
