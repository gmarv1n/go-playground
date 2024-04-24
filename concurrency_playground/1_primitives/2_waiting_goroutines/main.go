package main

import "sync"

func negativeWg() {
	wg := sync.WaitGroup{}
	wg.Add(-5)
	wg.Wait()
}

func zeroWg() {
	wg := sync.WaitGroup{}
	wg.Wait()
}

func main() {
	//negativeWg() // panic: sync: negative WaitGroup counter

	zeroWg() // just done
}
