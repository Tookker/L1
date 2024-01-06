package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
)

//9.Разработать конвейер чисел.
//Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

var wg sync.WaitGroup

func main() {
	const (
		max = 10
		min = 2
	)

	arrSize := rand.Intn(max-min) + max

	arr := make([]int, arrSize)
	for i := 0; i < arrSize; i++ {
		arr[i] = rand.Intn(max)
	}

	firstChan := make(chan int)
	secondChan := make(chan int)

	wg.Add(3)

	//Осуществляется запись в канал из массива
	go func(arr []int, ch chan<- int) {
		defer wg.Done()
		defer close(ch)

		for _, val := range arr {
			fmt.Println("Write to first chan value = ", val)
			ch <- val
		}
	}(arr, firstChan)

	//Передача данных из 1 канала во второй
	go func(chF <-chan int, chS chan<- int) {
		defer wg.Done()
		defer close(chS)

		for val := range chF {
			chS <- val
		}
	}(firstChan, secondChan)

	//Получение данных из 1 канала во 2, вывод квадрат числа в std::cout
	go func(ch <-chan int) {
		defer wg.Done()

		for val := range ch {
			fmt.Println("Pow ^2 val =", math.Pow(float64(val), 2))
		}
	}(secondChan)

	wg.Wait()
}
