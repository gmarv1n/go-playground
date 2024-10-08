package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once

	onceBody := func() {
		fmt.Println("Only once")
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			fmt.Println("iteration occurred")
			once.Do(onceBody)
		}()
	}

	wg.Wait()
}

// Lesson #_ video, time: __:__:__
