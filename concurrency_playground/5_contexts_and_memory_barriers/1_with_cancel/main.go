package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func recieveWeather(ctx context.Context, result chan struct{}, idx int) {
	value := rand.Intn(5000)
	randomTime := time.Duration(value) * time.Millisecond

	timer := time.NewTimer(randomTime)
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Printf("finished: %d\n", idx)
		result <- struct{}{}
	case <-ctx.Done():
		fmt.Printf("cancelled: %d\n", idx)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)

	ctx, cancel := context.WithCancel(context.Background())

	result := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(idx int) {
			defer wg.Done()
			recieveWeather(ctx, result, idx)
		}(i)
	}

	<-result
	cancel()

	wg.Wait()
}

// Lesson #7 video, time: 00:00:00
