package main

import (
	"fmt"
	"strings"
)

//20.Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func swapWords(str string) string {
	strArr := strings.Split(str, " ")

	for i := 0; i <= (len(strArr)-1)/2; i++ {
		strArr[i], strArr[len(strArr)-1-i] = strArr[len(strArr)-1-i], strArr[i]
	}

	return strings.Join(strArr, " ")
}

func main() {
	str := "snow dog sun"
	fmt.Println("Result =", swapWords(str))
}
