package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//6.Реализовать все возможные способы остановки выполнения горутины.

// остановка горутины через return
func returnStop() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("func returnStop fihished.")
		for {
			//TODO somefing...
			return
		}
	}()

	wg.Wait()
}

// остановка горутины через канал
func chanStop() {
	defer fmt.Println("func chanStop fihished.")

	chanWorker := make(chan int)
	chanExit := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	go func(worker chan<- int, exit <-chan struct{}) {
		defer wg.Done()
		worker <- rand.Int()
		<-exit
	}(chanWorker, chanExit)

	go func(chanExit chan<- struct{}) {
		defer wg.Done()
		time.Sleep(time.Second * 1)
		chanExit <- struct{}{}
	}(chanExit)

	fmt.Println("get from gorrutine", <-chanWorker)

	wg.Wait()
}

// остановка горутины через отмену контекста

func contextCancelStop() {
	defer fmt.Println("func contextCancelStop fihished.")

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)

	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				//TODO something...
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	cancel()

	wg.Wait()
}

// остановка горутины через таймаут контекста
func contextTimeOutStop() {
	defer fmt.Println("func contextTimeOutStop fihished.")

	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

	var wg sync.WaitGroup
	wg.Add(1)

	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				//TODO something...
			}
		}
	}(ctx)

	time.Sleep(3 * time.Second)
	wg.Wait()
}

// остановка горутины через закрытие канала
func contextCloseChan() {
	defer fmt.Println("func contextCloseChan fihished.")
	var localWg sync.WaitGroup

	mainChan := make(chan int)
	localWg.Add(1)
	go func(ch <-chan int) {
		defer localWg.Done()
		for i := range ch {
			//TODO something...
			i++
		}
	}(mainChan)

	for i := 0; i < 6; i++ {
		mainChan <- i
	}

	close(mainChan)
	localWg.Wait()
}

func main() {
	returnStop()
	chanStop()
	contextCancelStop()
	contextTimeOutStop()
	contextCloseChan()
}
