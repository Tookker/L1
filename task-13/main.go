package main

import "fmt"

//13. Поменять местами два числа без создания временной переменной.

func main() {

	val1, val2 := 10, 0
	fmt.Println(val1, val2)

	val1, val2 = val2, val1
	fmt.Println(val1, val2)
}
