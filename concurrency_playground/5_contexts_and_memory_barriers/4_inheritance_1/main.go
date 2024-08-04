package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// context.Background() - 1 уровень контекста
	// ctx - 2 уровень контекста с таймером на 2 секунды

	defer cancel()

	makeRequest(ctx)
}

func makeRequest(ctx context.Context) {
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()

	newCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	// newCtx - 3 уровень контекста с таймером на 10 сек
	defer cancel()

	select {
	case <-newCtx.Done():
		// сработает именно этот кейс, так как
		// отработает таймаут у parent контекста с 2 секундами
		// и завершит и дочерний
		fmt.Println("canceled")
	case <-timer.C:
		fmt.Println("timer")
	}
}

// Lesson #7 video, time: 00:__:__
