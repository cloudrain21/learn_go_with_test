package store_test

import (
	"github.com/cloudrain21/learn_go_with_test/mytestapp/server/store"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFileSystem(t *testing.T) {
	t.Run("stub store", func(t *testing.T) {
		filename := "/tmp/test.json"

		dbfile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
		assert.Equal(t, nil, err)

		sto, _ := store.NewFileSystemPlayerStore(dbfile)
		sto.PostPlayerScore("rain")

		got := sto.GetPlayerScore("rain")
		want := 1

		assert.Equal(t, want, got)
	})
}
