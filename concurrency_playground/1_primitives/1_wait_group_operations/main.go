package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func() {
			defer wg.Done()

			fmt.Println("test")
		}()
	}

	wg.Wait()
}
