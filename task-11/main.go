package main

import "fmt"

//11.Реализовать пересечение двух неупорядоченных множеств.

func main() {
	firstArr := []int{1, 2, 3, 4, 5}
	secondArr := []int{4, 5, 6, 7, 8, 1, 2}

	firstArrSize := len(firstArr)
	secondArrSize := len(secondArr)

	var maxLen int

	if firstArrSize >= secondArrSize {
		maxLen = firstArrSize
	} else {
		maxLen = secondArrSize
	}

	mapT := make(map[int]struct{})
	res := make([]int, 0, maxLen)

	// Проверка на наличие ключа в мапе
	contains := func(m map[int]struct{}, val int) {
		_, ok := m[val]
		//если есть такой ключ записываем его значение в результирующий слайс
		if ok {
			res = append(res, val)
		}
		m[val] = struct{}{}
	}

	// Пробегаемся по всем элементам двух сласов
	for i, j := 0, 0; i < firstArrSize || j < secondArrSize; i, j = i+1, j+1 {
		//Проверка на выход за пределы слайсы
		if i < firstArrSize {
			contains(mapT, firstArr[i])
		}

		//Проверка на выход за пределы слайсы
		if j < secondArrSize {
			contains(mapT, secondArr[j])
		}
	}

	fmt.Println("Result := ", res)
}
