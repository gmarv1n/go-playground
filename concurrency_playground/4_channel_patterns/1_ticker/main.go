package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	time.Sleep(5 * time.Second)

	ch <- 1
}

func main() {
	ch := make(chan int)
	go producer(ch)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case value := <-ch:
			fmt.Println(value)
			return
		case <-ticker.C:
			fmt.Println("tick")
		}
	}
}

// Lesson #6 video, time: 00:00:00
