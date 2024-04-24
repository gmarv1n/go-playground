package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mutex1 sync.Mutex
var mutex2 sync.Mutex

func goroutine1() {
	mutex1.Lock()

	runtime.Gosched()
	// Goshed - говорит планировщику что потуши меня и поставь
	// на выполнение какую-нибудь другую горутину
	// в итоге в этом кейсе выполнение уходит в горутин2
	// там лочится mutex2 и получается что обе горутины пытаются залочить
	// mutex2 и mutex1, а они уже все залочены. Ну и все
	for !mutex2.TryLock() {
		// active waiting
	}

	mutex2.Unlock()
	mutex1.Unlock()

	fmt.Println("goroutine1 finished")
}

func goroutine2() {
	mutex2.Lock()

	runtime.Gosched()
	for !mutex1.TryLock() {
		// active waiting
	}

	mutex1.Unlock()
	mutex2.Unlock()

	fmt.Println("goroutine2 finished")
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		goroutine1()
	}()

	go func() {
		defer wg.Done()
		goroutine2()
	}()

	wg.Wait()
}

// Lesson #_ video, time: __:__:__
