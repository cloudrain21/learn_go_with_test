package main

type InMemoryPlayerStore struct {
	score  map[string]int
	league League
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.score[name]
}

func (i *InMemoryPlayerStore) PostPlayerScore(name string) {
	i.score[name]++
}

func (i *InMemoryPlayerStore) GetLeagueTable() League {
	for name, wins := range i.score {
		i.league = append(i.league, Player{name, wins})
	}
	return i.league
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}, League{}}
}
