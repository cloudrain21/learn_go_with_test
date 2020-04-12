package store

import (
	"encoding/json"
	"errors"
	"io"
)

var (
	ErrJsonParse = errors.New("json parsing error")
)

type League []Player

func NewLeague(rdr io.Reader) (League, error) {
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		return nil, ErrJsonParse
	}

	return league, err
}

func (l League) find(name string) *Player {
	for i, p := range l {
		if name == p.Name {
			return &l[i]
		}
	}
	return nil
}
