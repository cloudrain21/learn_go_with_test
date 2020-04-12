package store_test

import (
	"github.com/cloudrain21/learn_go_with_test/mytestapp/server/store"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStub(t *testing.T) {
	t.Run("stub store", func(t *testing.T) {
		sto := store.NewStubPlayerStore()
		sto.PostPlayerScore("rain")

		got := sto.GetPlayerScore("rain")
		want := 1

		assert.Equal(t, want, got)
	})
}
