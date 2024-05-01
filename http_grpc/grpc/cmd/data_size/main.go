package main

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/gmarv1n/go-playground/http_grpc/pkg/note_v1"
	"google.golang.org/protobuf/proto"
)

func main() {
	session := &note_v1.NoteInfo{
		Title:    gofakeit.BuzzWord(),
		Content:  gofakeit.IPv4Address(),
		Author:   gofakeit.BeerIbu(),
		IsPublic: gofakeit.Bool(),
	}

	dataJson, _ := json.Marshal(session)
	fmt.Printf("dataJson len %d byte \n%v\n", len(dataJson), dataJson)

	dataPb, _ := proto.Marshal(session)
	fmt.Printf("dataPb len %d byte \n%v\n", len(dataPb), dataPb)
}
