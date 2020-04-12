package store_test

import (
	"github.com/cloudrain21/learn_go_with_test/mytestapp/server/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemory(t *testing.T) {
	t.Run("inmemory store", func(t *testing.T) {
		sto := store.NewInMemoryPlayerStore()
		sto.PostPlayerScore("rain")

		got := sto.GetPlayerScore("rain")
		want := 1

		assert.Equal(t, want, got)
	})
}
