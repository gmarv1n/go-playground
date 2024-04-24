package main

import "sync"

func lockAnyTimes() {
	mu := sync.Mutex{}
	mu.Lock()
	mu.Lock()
}

func unlockWithoutLock() {
	mu := sync.Mutex{}
	mu.Unlock()
}

func unlockFromOtherGoroutine() {
	mu := sync.Mutex{}
	mu.Lock()

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		mu.Unlock()
	}()

	wg.Wait()

	mu.Lock()
	mu.Unlock()
}

func main() {
	// lockAnyTimes() // fatal error, deadlock cause of two locks

	// unlockWithoutLock() // fatal error: unlock of unlocked mu

	unlockFromOtherGoroutine() // all ok cause of wg.Wait() and no matter what goroutine it unlocks
}
