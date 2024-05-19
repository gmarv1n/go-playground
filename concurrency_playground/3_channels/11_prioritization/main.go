package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	for {
		ch <- 1
		time.Sleep(time.Second)
	}
}

func main() {
	ch1 := make(chan int) // этот приеоритетней
	ch2 := make(chan int)

	go producer(ch1)
	go producer(ch2)

	for {
		// ну тут типа прикол в том что первоначально проверяешь
		// более важный канал ну и выходишь по дефолту, а потом уже
		// типа ну как получится
		select {
		case value := <-ch1:
			fmt.Println(value)
		default:

		}

		// ну как получится это про вот это:
		select {
		case value := <-ch1:
			fmt.Println(value)
		case value := <-ch2:
			fmt.Println(value)
		}
	}
}

// Lesson #5 video, time: __:__:__
