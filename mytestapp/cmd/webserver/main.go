package main

import (
	"github.com/cloudrain21/learn_go_with_test/mytestapp/domain/server"
	"github.com/cloudrain21/learn_go_with_test/mytestapp/domain/server/store"
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

	store, _ := store.NewFileSystemPlayerStore(dbfile)
	server := server.NewPlayerServer(store)

	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Fatalf("error : %v", err)
	}
}
