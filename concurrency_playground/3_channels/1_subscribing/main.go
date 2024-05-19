package main

import (
	"fmt"
	"sync"
)

func notifier(signals chan int) {
	signals <- 1000 // это блокирующий вызов - горутина заблокируется, пока другая это не прочитает
	// ВАЖНО: это будет ВСЕГДА блокирующий вызов только для небуферизированного канала
	close(signals)
}

func subscriber(signals chan int) {
	value, ok := <-signals // это блокирующий вызов - горутина будет ждать значения
	fmt.Println(value, ok)
}

func main() {
	signals := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {
		defer wg.Done()
		notifier(signals)
	}()

	go func() {
		defer wg.Done()
		subscriber(signals)
	}()

	go func() {
		defer wg.Done()
		subscriber(signals)
	}()

	wg.Wait()
}

// Lesson #5 video, time: __:__:__
