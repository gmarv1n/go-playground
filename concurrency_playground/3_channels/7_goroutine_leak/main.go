package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(strings <-chan string) <-chan struct{} {
		completed := make(chan struct{})
		go func() {
			defer func() {
				// так как канал не закрывается - горутина и канал будут висеть
				// в памяти и ждать поэтому принта не будет
				//
				// вот это и есть утечка памяти
				fmt.Println("doWork exited")
				close(completed)
			}()

			for str := range strings {
				fmt.Println(str)
			}
		}()

		return completed // это вообще не нужно, просто зачем то есть
		// в примере
	}

	strings := make(chan string)
	doWork(strings)
	strings <- "Test"

	time.Sleep(time.Second)
	fmt.Println("Done")
}

// Lesson #5 video, time: __:__:__
