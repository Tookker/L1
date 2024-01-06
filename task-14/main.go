package main

import "fmt"

//14.Разработать программу,
// которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func main() {
	var v1 int
	var v2 string
	var v3 bool
	var v4 chan float32

	arr := [4]interface{}{v1, v2, v3, v4}

	for i := 0; i < 4; i++ {
		if _, ok := arr[i].(string); ok {
			fmt.Println("It string")
		} else if _, ok = arr[i].(int); ok {
			fmt.Println("It int")
		} else if _, ok = arr[i].(bool); ok {
			fmt.Println("It bool")
		} else if _, ok = arr[i].(chan float32); ok {
			fmt.Println("It chan float32")
		}
	}
}
