package main

import (
	"fmt"
	"sync"
)

func notifier(signals chan int) {
	close(signals) // когда канал закроется - он отдас zero val всем кто ждал
}

func subscriber(signals chan int) {
	<-signals // пока канал не закрыт - горутины будут ждать что в него придет
	fmt.Println("signalled")
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
