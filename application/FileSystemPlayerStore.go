package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() League {
	f.database.Seek(0, 0)
	league, _ := NewLeague(f.database)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	p := f.GetLeague().Find(name)
	if p == nil {
		return 0
	}
	return p.Wins
}

func (f *FileSystemPlayerStore) PostPlayerScore(name string) {
	league := f.GetLeague()

	player := league.Find(name)
	if player == nil {
		league = append(league, Player{name, 0})
		player = &league[len(league)-1]
	}
	player.Wins++

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(league)
}
