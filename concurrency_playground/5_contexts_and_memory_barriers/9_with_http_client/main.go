package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		1*time.Millisecond,
	)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://example.com",
		nil,
	)

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}
}

// Lesson #7 video, time: 00:28:53
