package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	filename := "/tmp/test.json"

	dbfile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("open file error : %s : %v\n", filename, err)
	}

	store := &FileSystemPlayerStore{dbfile}

	server := NewPlayerServer(store)
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("error : %v", err)
	}
}
