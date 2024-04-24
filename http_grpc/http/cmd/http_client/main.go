package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"github.com/fatih/color"
	"github.com/go-chi/chi"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const (
	baseUrl       = "http://localhost:8081"
	createPostfix = "/notes"
	getPostfix    = "/notes/%d"
)

type NoteInfo struct {
	Title    string `json:"title"`
	Context  string `json:"context"`
	Author   string `json:"author"`
	IsPublic bool   `json:"is_public"`
}

type Note struct {
	ID        int64     `json:"id"`
	Info      *NoteInfo `json:"info"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SyncMap struct {
	elems map[int64]*Note
	m     sync.RWMutex
}

var notes = &SyncMap{
	elems: make(map[int64]*Note),
}

func createNoteHandler(w http.ResponseWriter, r *http.Request) {
	info := &NoteInfo{}

	if err := json.NewDecoder(r.Body); err != nil {
		http.Error(w, "failed to decode note data", http.StatusBadRequest)
		return
	}

	rand.Seed(time.Now().UnixNano())
	now := time.Now()

	note := &Note{
		ID:        rand.Int63(),
		Info:      info,
		CreatedAt: now,
		UpdatedAt: now,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "failed to encode note data", http.StatusInternalServerError)
		return
	}

	notes.m.Lock()
	defer notes.m.Unlock()

	notes.elems[note.ID] = note
}

func getNoteHandler(w http.ResponseWriter, r *http.Request) {
	noteID := chi.URLParam(r, "id")
	id, err := parseNoteID(noteID)
	if err != nil {
		http.Error(w, "failed to parse note ID", http.StatusBadRequest)
		return
	}

	notes.m.RLock()
	defer notes.m.RUnlock()

	note, ok := notes.elems[id]
	if !ok {
		http.Error(w, "Note not found", http.StatusNotFound)
		return
	}

	if err = json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "failed to encode note data", http.StatusInternalServerError)
		return
	}
}

func parseNoteID(noteID string) (int64, error) {
	id, err := strconv.ParseInt(noteID, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func createNoteClient() (Note, error) {
	note := &NoteInfo{
		Title:    gofakeit.CarMaker(),
		Context:  gofakeit.BeerAlcohol(),
		Author:   gofakeit.BuzzWord(),
		IsPublic: gofakeit.Bool(),
	}

	data, err := json.Marshal(note)
	if err != nil {
		return Note{}, err
	}

	resp, err := http.Post(baseUrl+createPostfix, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return Note{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return Note{}, errors.New(color.RedString("Get note failed: not 201 status code\n"))
	}

	var createdNote Note

	if err := json.NewDecoder(resp.Body).Decode(&createdNote); err != nil {
		return Note{}, err
	}

	return createdNote, nil
}

func getNoteClient(id int64) (Note, error) {
	resp, err := http.Get(fmt.Sprintf(baseUrl+getPostfix, id))
	if err != nil {
		return Note{}, err
	}

	if resp.StatusCode != http.StatusOK {

	}

	var note Note

	if err := json.NewDecoder(resp.Body).Decode(&note); err != nil {
		return Note{}, errors.New(color.RedString("Get note failed: not 201 status code\n"))
	}

	return note, nil
}

func main() {
	note, err := createNoteClient()
	if err != nil {
		log.Fatal("failed to create note:", err)
	}

	log.Printf(color.RedString("Note created:\n"), color.GreenString("%+v", note))

	recievedNote, err := getNoteClient(note.ID)
	if err != nil {
		log.Fatal("failed to get note: ", err)
	}

	log.Printf(color.RedString("Note info got:\n"), color.GreenString("%+v", recievedNote))
}
