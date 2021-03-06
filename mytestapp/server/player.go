package server

import (
	"encoding/json"
	"fmt"
	"github.com/cloudrain21/learn_go_with_test/mytestapp/server/store"
	"net/http"
	"path"
)


type PlayerServer struct {
	store store.PlayerStore
	http.Handler
}

//func (p *PlayerServer)ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	p.Handler.ServeHTTP(w,r)
//}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(p.store.GetLeagueTable())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.showScore(w, r)
	case http.MethodPost:
		p.postScore(w, r)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	name := path.Base(r.URL.Path)
	if p.store.GetPlayerScore(name) == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, p.store.GetPlayerScore(name))
}

func (p *PlayerServer) postScore(w http.ResponseWriter, r *http.Request) {
	name := path.Base(r.URL.Path)
	w.WriteHeader(http.StatusAccepted)
	p.store.PostPlayerScore(name)
}

func NewPlayerServer(store store.PlayerStore) *PlayerServer {
	p := new(PlayerServer)

	p.store = store
	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playerHandler))

	p.Handler = router

	return p
}
