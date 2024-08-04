package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "hello world\n")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Print(err.Error())
		}
	}()

	<-ctx.Done()

	timeOutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeOutCtx); err != nil {
		log.Print(err.Error())
	}

	fmt.Println("cancelled")
}

// Lesson #7 video, time: 00:35:26
