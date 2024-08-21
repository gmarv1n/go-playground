package main

import (
	"context"
	"fmt"
	"time"
)

func WithAfterFunc(
	ctx context.Context,
	action func(),
) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())

	if action != nil {
		go func() {
			<-ctx.Done()
			action()
		}()
	}

	return ctx, cancel
}

func main() {
	afterDone := func() {
		// some wotk
		fmt.Println("done after ctx")
	}

	_, cancel := WithAfterFunc(context.Background(), afterDone)
	cancel()

	time.Sleep(1 * time.Second)
}

// Lesson #7 video, time: 00:35:26
