package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//18.Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

// Счетчик через мьютексы
type mutexInc struct {
	mutex   sync.Mutex
	Counter int
}

func (m *mutexInc) inc() {
	defer m.mutex.Unlock()
	m.mutex.Lock()
	m.Counter++
}

// Счетчик черех атомик
type atomicInc struct {
	Counter int64
}

func (a *atomicInc) inc() {
	atomic.AddInt64(&a.Counter, 1)
}

func main() {
	var wg sync.WaitGroup
	var mutInc mutexInc

	wg.Add(15)
	for i := 0; i < 15; i++ {
		go func() {
			defer wg.Done()
			mutInc.inc()
		}()

	}

	wg.Wait()

	fmt.Println("Mutex inc =", mutInc.Counter)

	var atomInc atomicInc
	wg.Add(20)
	for i := 0; i < 20; i++ {
		go func() {
			defer wg.Done()
			atomInc.inc()
		}()

	}
	wg.Wait()

	fmt.Println("Atomic inc =", atomInc.Counter)
}
