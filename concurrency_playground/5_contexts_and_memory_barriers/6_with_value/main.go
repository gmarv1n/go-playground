package main

import (
	"context"
	"fmt"
)

func main() {
	traceCtx := context.WithValue(
		context.Background(),
		"trace_id",
		"31-07-2024",
	)

	traceCtx2 := context.WithValue(
		traceCtx,
		"foo",
		"bar",
	)

	makeRequest(traceCtx)
	makeRequest(traceCtx2)
}

func makeRequest(ctx context.Context) {
	oldValue, ok := ctx.Value("trace_id").(string)
	if ok {
		fmt.Println(oldValue)
	}

	newCtx := context.WithValue(
		ctx,
		"trace_id",
		"01-08-2024",
	)
	newValue, ok := newCtx.Value("trace_id").(string)
	if ok {
		fmt.Println(newValue)
		// у дочернего контекста значение по ключу переопределится
	}

	foo, ok := newCtx.Value("foo").(string)
	if ok {
		fmt.Println(foo)
		// а вот значения из parent контекстов сохранятся
	}
}

// Lesson #7 video, time: 00:__:__
