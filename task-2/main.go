package main

import (
	"fmt"
	"math"
	"sync"
)

//2.Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10)
// и выведет их квадраты в stdout.

var wg sync.WaitGroup

const pow = 2.0

// Возведение числа в квадрат и вывод в stdout
func getPowNum(val int) {
	fmt.Printf("Pow of %d = %f\n", val, math.Pow(float64(val), pow))
	//Уменьшение счетчика WaitGroup
	wg.Done()
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	lenArr := len(arr)
	//Добавление в счетчик wg ожидаемое количество выполнямых потоков
	wg.Add(lenArr)

	for i := 0; i < lenArr; i++ {
		go getPowNum(arr[i])
	}

	//Остановка основного потока до тех пор пока счетчик wg не будет = 0
	wg.Wait()
}
