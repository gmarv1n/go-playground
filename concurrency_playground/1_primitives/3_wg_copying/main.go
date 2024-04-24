package main

import "sync"

func doneCopy(wg sync.WaitGroup) {
	wg.Done()
}

func donePointer(wg *sync.WaitGroup) {
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)

	//doneCopy(wg) // fatal error: all goroutines are asleep - deadlock!

	donePointer(&wg) // ok

	wg.Wait()
}
