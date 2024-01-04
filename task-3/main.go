package main

import (
	"fmt"
	"math"
	"sync"
)

//3.Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

var wg sync.WaitGroup
var mutex sync.Mutex
var res int64

const pow = 2.0

// Возведение числа в квадрат
// При входе в функцию лочим мьютекс, по завершении функции открываем мьютекс.
func powNum(val int) {
	mutex.Lock()
	defer mutex.Unlock()

	res = res + int64(math.Pow(float64(val), pow))

	wg.Done()
}

func main() {
	arr := []int{2, 4, 6, 8, 10}
	lenArr := len(arr)
	wg.Add(lenArr)

	for i := 0; i < lenArr; i++ {
		go powNum(arr[i])
	}

	wg.Wait()
	fmt.Printf("Res = %d\n", res)
}
