package main

import (
	"context"
	"fmt"
)

func main() {
	traceCtx := context.WithValue(
		context.Background(),
		"trace_id",
		"321",
	)

	makeRequest(traceCtx)
}

func makeRequest(ctx context.Context) {
	oldValue, ok := ctx.Value("trace_id").(string)
	if ok {
		fmt.Println(oldValue)
	}

	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	newValue, ok := newCtx.Value("trace_id").(string)
	if ok {
		fmt.Println(newValue)
		// у дочернего контекста значение наследуется
	}
}

// Lesson #7 video, time: 00:__:__
