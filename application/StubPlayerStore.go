package main

type StubPlayerStore struct {
	score  map[string]int
	league League
}

func (i *StubPlayerStore) GetPlayerScore(name string) int {
	return i.score[name]
}

func (i *StubPlayerStore) PostPlayerScore(name string) {
	i.score[name]++
}

func (s *StubPlayerStore) GetLeagueTable() League {
	return s.league
}

func NewStubPlayerStore() *StubPlayerStore {
	return &StubPlayerStore{map[string]int{}, League{}}
}
