package main

import (
	"fmt"
	"sync"
)

func race() {
	text := ""

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		text = "hello race"
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Ex #1", text)
	}()

	wg.Wait()
}

func noRace() {
	text := ""
	mu := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer func() {
			wg.Done()
			mu.Unlock()
		}()

		mu.Lock()
		text = "hello race"
	}()

	go func() {
		defer func() {
			wg.Done()
			mu.Unlock()
		}()

		mu.Lock()
		fmt.Println("Ex #2", text)
	}()

	wg.Wait()
}

func main() {
	race()   // go run -race pool_test.go - ран с рейс детектором
	noRace() // No data race here, but race condition
}

// Lesson #3 video, time: 00:32:24
