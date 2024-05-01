package main

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/gmarv1n/go-playground/http_grpc/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

const grpcPort = 50051

type server struct {
	note_v1.UnimplementedNoteV1Server
}

func (s *server) Get(ctx context.Context, req *note_v1.GetRequest) (*note_v1.GetResponse, error) {
	log.Printf("Note id: %d", req.GetId())
	return &note_v1.GetResponse{
		Note: &note_v1.Note{
			Id: req.GetId(),
			Note: &note_v1.NoteInfo{
				Title:    gofakeit.BuzzWord(),
				Content:  gofakeit.BuzzWord(),
				Author:   gofakeit.BuzzWord(),
				IsPublic: gofakeit.Bool(),
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatal("failed to listen server: %w", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	note_v1.RegisterNoteV1Server(s, &server{})

	log.Printf("server listening on %s", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatal("failed to serve: %w", err)
	}
}
