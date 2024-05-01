package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1000)

	var value atomic.Int32
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			value.Add(1)
			// тут важно что именно Add, через store(load)
			// получится тот же рейс кондишн
		}()
	}

	wg.Wait()

	fmt.Println(value.Load())
}
