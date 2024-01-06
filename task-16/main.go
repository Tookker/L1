package main

import "fmt"

//16.Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func quicksort(arr []int) []int {
	const min = 1

	if len(arr) <= min {
		return arr
	}

	startVal := arr[0]
	var lower, higher []int

	//распределяем меньшие либо равные значения чем arr[0] в слайс lower, в higher большие
	for _, val := range arr[1:] {
		if val <= startVal {
			lower = append(lower, val)
		} else {
			higher = append(higher, val)
		}
	}

	//Вызываем рекурсивно фунцию для lower и higher слайсов
	lower = quicksort(lower)
	higher = quicksort(higher)

	return append(append(lower, startVal), higher...)
}

func main() {
	arr := []int{5, 8, 2, 1, 7, 9, 3, 1, 4, 6, 9, 0, 33, 1}

	fmt.Println("Result = ", quicksort(arr))
}
