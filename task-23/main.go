package main

import (
	"errors"
	"fmt"
)

//23.Удалить i-ый элемент из слайса.

// Способ с сохранением целостности данных входящего массива
func removeIndxSliseSave[T comparable](indx uint, arr []T) ([]T, error) {
	var arrSize = len(arr)

	if arrSize <= int(indx) {
		return []T{}, errors.New("Invalid index")
	}

	newArr := make([]T, arrSize)
	copy(newArr, arr)

	if indx == 0 {
		return newArr[1:], nil
	}

	if int(indx) == arrSize-1 {
		return newArr[:arrSize-1], nil
	}

	return append(newArr[:indx], newArr[indx+1:]...), nil
}

// Способ без сохранения целостности данных входящего массива
func removeIndxSlise[T comparable](indx uint, arr []T) ([]T, error) {
	var arrSize = len(arr)

	if arrSize <= int(indx) {
		return []T{}, errors.New("Invalid index")
	}

	if indx == 0 {
		return arr[1:], nil
	}

	if int(indx) == arrSize-1 {
		return arr[:arrSize-1], nil
	}

	return append(arr[:indx], arr[indx+1:]...), nil
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(removeIndxSliseSave[int](0, arr))
	fmt.Println(removeIndxSliseSave[int](9, arr))
	fmt.Println(removeIndxSliseSave[int](5, arr))
	fmt.Println(arr)
	fmt.Println("========================")

	fmt.Println(removeIndxSlise[int](0, arr))
	fmt.Println(removeIndxSlise[int](9, arr))
	fmt.Println(removeIndxSlise[int](5, arr))
	fmt.Println(arr)
}
