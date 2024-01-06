package main

import (
	"errors"
	"fmt"
)

//17.Реализовать бинарный поиск встроенными методами языка.

func binSearch(val int, arr []int) (int, error) {
	const div = 2
	arrLen := len(arr)

	if arrLen == 0 {
		return 0, errors.New("Slice is empty")
	}

	if arrLen == 1 {
		return arr[0], nil
	}

	var center int
	min, max := 0, arrLen-1

	for min <= max {
		center = (min + max) / 2

		if arr[center] == val {
			return center, nil
		} else if arr[center] < val {
			min = center + 1
		} else {
			max = center - 1
		}
	}

	return 0, errors.New("Element not found")
}

func main() {
	arr := []int{1, 5, 7, 9, 12}

	res, err := binSearch(9, arr)

	fmt.Println("Res =", res, err)
}
