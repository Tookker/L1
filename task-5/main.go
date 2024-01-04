package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

//5.Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.

func main() {
	workerCount := flag.Int("t", 6, "Timeout")
	flag.Parse()
	if *workerCount <= 0 {
		log.Fatalln("Err input")
	}

	wg.Add(2)

	workerChan := make(chan int)
	//Инициализация контекста с таймаутом
	ctx, _ := context.WithTimeout(context.Background(), time.Second*time.Duration(*workerCount))
	go func(c chan<- int, ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				close(c)
				fmt.Println("Close chan")
				return
			default:
				c <- rand.Int()
				fmt.Println("Write to chan...")
			}
		}
	}(workerChan, ctx)

	go func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			fmt.Println("Read from chan ", i)
		}
	}(workerChan)

	wg.Wait()
}
