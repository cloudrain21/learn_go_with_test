package store

import (
	"encoding/json"
	"io"
	"sort"
)

type FileSystemPlayerStore struct {
	database io.ReadWriteSeeker
	league   League
}

func NewFileSystemPlayerStore(d io.ReadWriteSeeker) (*FileSystemPlayerStore, error) {
	d.Seek(0,0)
	league, err := NewLeague(d)
	if err != nil {
		return nil, err
	}
	return &FileSystemPlayerStore{d, league}, nil
}

func (f *FileSystemPlayerStore) GetLeagueTable() League {
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins > f.league[j].Wins
	})
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	p := f.GetLeagueTable().Find(name)
	if p == nil {
		return 0
	}
	return p.Wins
}

func (f *FileSystemPlayerStore) PostPlayerScore(name string) {
	player := f.league.Find(name)
	if player == nil {
		f.league = append(f.league, Player{name, 0})
		player = &f.league[len(f.league)-1]
	}
	player.Wins++

	f.database.Seek(0, 0)
	json.NewEncoder(f.database).Encode(f.league)
}
