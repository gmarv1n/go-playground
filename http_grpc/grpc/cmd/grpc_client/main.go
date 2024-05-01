package main

import (
	"context"
	"github.com/fatih/color"
	"github.com/gmarv1n/go-playground/http_grpc/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	address = "localhost:50051"
	noteID  = 12
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect to grpc server: %w", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal("failed to close  grpc connection: %w", err)
		}
	}()

	c := note_v1.NewNoteV1Client(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Get(ctx, &note_v1.GetRequest{Id: noteID})
	if err != nil {
		log.Fatal("failed to get note by id: %w", err)
	}

	log.Printf(color.RedString("Note info: \n"), color.GreenString("%+v", r.Note))
}
