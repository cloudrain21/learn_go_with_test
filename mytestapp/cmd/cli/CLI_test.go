package main

import (
	"github.com/cloudrain21/learn_go_with_test/mytestapp/domain/server/store"
	"testing"
)

type CLI struct {
	playStore store.PlayerStore
}

func (c *CLI) PlayPoker() {
	c.playStore.PostPlayerScore("Cleo")
}

func TestCLI(t *testing.T) {
	playerStore := store.NewStubPlayerStore()
	cli := &CLI{playerStore}
	cli.PlayPoker()

	if playerStore.GetPlayerScore("Cleo") != 1 {
		t.Fatalf("expected a win call but didn't get any : %d\n", len(playerStore.GetLeagueTable()))
	}
}
