package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

//4.Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

var wg sync.WaitGroup

// woker читает данные из канала
func woker(id int, a <-chan int) {
	defer wg.Done()
	for j := range a {
		fmt.Println("Worker", id, "processed job:", j)
	}
}

func main() {
	workerCount := flag.Int("c", 1, "Count of workers")
	//чтение аргументов
	flag.Parse()

	if *workerCount <= 0 {
		log.Fatalln("Err input")
	}

	//иницилизая контекста с отменой
	ctx, cancel := context.WithCancel(context.Background())
	//иницилизация канала в который будем записывать данные
	mainChan := make(chan int)
	wg.Add(*workerCount)

	for i := 0; i < *workerCount; i++ {
		go woker(i, mainChan)
	}

	//запуск цикла, условием выхода является отмена контекста
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(mainChan)
				return
			default:
				mainChan <- rand.Int()
			}
		}
	}()

	//инициализация канал которйы принимает сигналы от ОС
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Get signal from system", <-c, "Stoping...")

	//установить контекст в отмену
	cancel()
	wg.Wait()

	fmt.Println("All gorrutines stopped")
}
