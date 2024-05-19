package main

import "sync"

func notifier(signals chan<- struct{}) {
	signals <- struct{}{} // тут пустая структура и типа это будет 0 байт
}

func subscriber(signals <-chan struct{}) {
	<-signals
}

func main() {
	signals := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		notifier(signals)
	}()

	go func() {
		defer wg.Done()
		subscriber(signals)
	}()

	wg.Wait()
}

// Lesson #5 video, time: __:__:__
