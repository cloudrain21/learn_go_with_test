package store_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/cloudrain21/learn_go_with_test/mytestapp/server/store"
)

func TestInMemory(t *testing.T) {
	t.Run("stub store", func(t *testing.T) {
		sto := store.NewInMemoryPlayerStore()
		sto.PostPlayerScore("rain")

		got := sto.GetPlayerScore("rain")
		want := 1

		assert.Equal(t, want, got)
	})
}
