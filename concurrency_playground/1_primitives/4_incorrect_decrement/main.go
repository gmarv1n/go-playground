package main

import (
	"fmt"
	"sync"
)

// эта хуйня называется "Проблема потерянного обновления"
func incorrect() {
	wg := sync.WaitGroup{}
	wg.Add(1000)

	value := 0
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			value++
		}()
	}

	wg.Wait()
	fmt.Println(value)
}

func correct() {
	wg := sync.WaitGroup{}
	wg.Add(1000)

	mu := sync.Mutex{} // Mutex = Mutual Exclusion, Взаимное исключение

	value := 0
	for i := 0; i < 1000; i++ {
		go func() {
			defer func() {
				wg.Done()
				mu.Unlock()
			}()

			mu.Lock()
			value++
		}()
	}

	wg.Wait()
	fmt.Println(value)
}

func main() {
	incorrect()

	correct()
}
