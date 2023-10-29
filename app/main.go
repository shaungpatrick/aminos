package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type Aminos struct {
	Aimnos []Amino `json:"aminos"`
}

type Amino struct {
	Name       string   `json:"name"`
	Compressed string   `json:"compressed"`
	Codons     []string `json:"codons"`
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		w.Write([]byte("err"))
	}
}

func main() {
	r := chi.NewRouter()
	r.Get("/", getAllAimnos)
	http.ListenAndServe(":3333", r)
}

func getAllAimnos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	a := getAminosFromFile()
	json.NewEncoder(w).Encode(a)
}

func getAmino(name string) (*Amino, error) {
	aminos := getAminosFromFile()
	for _, a := range aminos.Aimnos {
		if name == a.Name {
			return &a, nil
		}
	}
	return nil, errors.New("Amino not found.")
}

func getAminosFromFile() *Aminos {
	data, err := os.ReadFile("./aminos.json")
	if err != nil {
		fmt.Print(err)
	}
	a := Aminos{}
	err = json.Unmarshal(data, &a)
	return &a
}
