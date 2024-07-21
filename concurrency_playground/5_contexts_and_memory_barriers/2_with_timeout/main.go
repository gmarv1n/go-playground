package main

import (
	"context"
	"fmt"
	"time"
)

func makeRequest(ctx context.Context) {
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()

	select {
	case <-timer.C:
		fmt.Println("finished")
	case <-ctx.Done():
		fmt.Println("cancelled")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel() // если WithCancel или WithTimeOut не будет отменен - будет
	// утчека памяти в горутине

	makeRequest(ctx)
}

// Lesson #7 video, time: 00:00:00
