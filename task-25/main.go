package main

import (
	"context"
	"sync"
	"time"
)

//25.Реализовать собственную функцию sleep.

func mySleepMs(t uint) {
	ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(t))
	var wg sync.WaitGroup

	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				{

				}
			}
		}
	}(ctx, &wg)

	wg.Wait()
}

func main() {
	mySleepMs(5000)
}
