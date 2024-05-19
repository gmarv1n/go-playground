package main

import (
	"fmt"
	"sync"
)

var ch chan int

func producer() {
	for i := 0; i < 5; i++ {
		ch <- i // на этом месте горутина заблочится до того как
		// консумер вычитает значение
	}

	close(ch)
}

func consumer() {
	// range по каналу будет работать до тех пор, пока не закроется канал
	// так же на каждой итерации будет блокироваться горутина
	for value := range ch {
		fmt.Println(value)
	}
}

func main() {
	ch = make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		producer()
	}()

	go func() {
		defer wg.Done()
		consumer()
	}()

	wg.Wait()
}

// Lesson #5 video, time: __:__:__
