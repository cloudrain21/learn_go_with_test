package store

import "sort"

type StubPlayerStore struct {
	score  map[string]int
	league League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	return s.score[name]
}

func (s *StubPlayerStore) PostPlayerScore(name string) {
	s.score[name]++
}

func (s *StubPlayerStore)findLeaguePlayerPos(name string) int {
	for idx, p := range s.league {
		if name == p.Name {
			return idx
		}
	}
	return -1
}

func (s *StubPlayerStore) GetLeagueTable() League {
	for name, wins := range s.score {
		idx := s.findLeaguePlayerPos(name)
		if idx >= 0 {
			s.league[idx].Wins = wins
		} else {
			s.league = append(s.league, Player{name,wins})
		}
	}

	sort.Slice(s.league, func(i, j int) bool {
		return s.league[i].Wins > s.league[j].Wins
	})
	return s.league
}

func NewStubPlayerStore() *StubPlayerStore {
	return &StubPlayerStore{map[string]int{}, League{}}
}
