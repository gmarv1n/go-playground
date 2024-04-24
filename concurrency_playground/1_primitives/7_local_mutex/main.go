package main

import (
	"fmt"
	"sync"
)

var value int
var valueOk int

func incWithLocalMu() {
	mu := sync.Mutex{}

	mu.Lock()
	value++
	mu.Unlock()
}

func inc(mu *sync.Mutex) {
	mu.Lock()
	valueOk++
	mu.Unlock()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			incWithLocalMu()
		}()
	}

	wg.Wait()

	fmt.Println(value) // not determ cause of local mu

	mu := &sync.Mutex{}
	wg2 := sync.WaitGroup{}
	wg2.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg2.Done()
			inc(mu)
		}()
	}

	wg2.Wait()

	fmt.Println(valueOk) // determ
}
