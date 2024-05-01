package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var data map[string]string
var initialized atomic.Bool

func initialize() {
	//if !initialized.Load() {
	//	runtime.Gosched()
	//	// в этом месте может заломитсья другая гоуртина,
	//	// кеторая до этого прервалась и будет вызов store(true)
	//	// несколько раз
	//	initialized.Store(true)
	//	data = make(map[string]string)
	//	fmt.Println("initialized")
	//}

	// правильно в таких кейсах юзать:
	if initialized.CompareAndSwap(false, true) { // CAS (Compare And Swap)
		data = make(map[string]string)
		fmt.Println("initialized")
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			initialize()
		}()
	}

	wg.Wait()
}

// Lesson #_ video, time: __:__:__
