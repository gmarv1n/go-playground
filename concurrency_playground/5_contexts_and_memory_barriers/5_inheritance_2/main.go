package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// ctx - 1 уровень с таймаутом 5 сек
	defer cancel()

	_, cancel = context.WithCancel(ctx)
	// _ - 2 уровень с кенселом
	cancel()
	// сразу завершаем дочерний контекст

	select {
	case <-ctx.Done():
		// дочерний контекст завершается там себе
		// но родительский продолжает жить
		fmt.Println("canceled")
	}
}

// Lesson #7 video, time: 00:__:__
