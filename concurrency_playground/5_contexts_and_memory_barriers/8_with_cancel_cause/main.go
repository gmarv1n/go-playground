package main

import (
	"context"
	"errors"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancelCause(context.Background())
	cancel(errors.New("error"))

	if ctx.Err() != nil {
		err := context.Cause(ctx)
		fmt.Println(err.Error())
	}
}

// Lesson #7 video, time: 00:28:53
