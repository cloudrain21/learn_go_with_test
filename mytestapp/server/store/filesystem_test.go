package store_test

import (
	"github.com/cloudrain21/learn_go_with_test/mytestapp/server/store"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFileSystem(t *testing.T) {
	t.Run("filesystem store", func(t *testing.T) {
		filename := "/tmp/test.json"
		dbfile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
		assert.Equal(t, nil, err)

		sto := store.CreateFileSystemPlayerStore(dbfile)
		assert.Equal(t, nil, err)

		sto.PostPlayerScore("rain")
		sto.PostPlayerScore("rain")

		got := sto.GetPlayerScore("rain")
		want := 2

		assert.Equal(t, want, got)
	})
}
