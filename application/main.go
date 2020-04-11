package main

import (
	"log"
	"net/http"
)

func main() {
	league := []Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}

	store := &InMemoryPlayerStore{nil, league}
	server := NewPlayerServer(store)
	if err := http.ListenAndServe(":9000", server); err != nil {
		log.Fatalf("error : %v", err)
	}
}
