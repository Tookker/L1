package main

import "fmt"

//19.Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
//Символы могут быть unicode.

func main() {
	str := "главрыба"

	runeSlise := []rune(str)
	for i := 0; i < (len(runeSlise)-1)/2; i++ {
		runeSlise[i], runeSlise[len(runeSlise)-1-i] = runeSlise[len(runeSlise)-1-i], runeSlise[i]
	}

	fmt.Println("Res", string(runeSlise))
}
