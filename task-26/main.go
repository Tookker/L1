package main

import (
	"errors"
	"fmt"
	"strings"
)

//26.Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func checkUniqeSim(str string) (bool, error) {
	var strSize = len(str)

	const errSizeZero = "str size = 0"

	if strSize == 0 {
		fmt.Println(errSizeZero)
		return false, errors.New(errSizeZero)
	}

	if strSize == 1 {
		return true, nil
	}

	newStr := strings.ToLower(str)
	runeArr := []rune(newStr)
	mapSim := make(map[rune]struct{})

	for _, val := range runeArr {
		_, ok := mapSim[val]
		if ok {
			return false, nil
		} else {
			mapSim[val] = struct{}{}
		}
	}

	return true, nil
}

func main() {

	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aAbcd"
	str4 := "aabcDd"

	fmt.Println(checkUniqeSim(str1))
	fmt.Println(checkUniqeSim(str2))
	fmt.Println(checkUniqeSim(str3))
	fmt.Println(checkUniqeSim(str4))
}
