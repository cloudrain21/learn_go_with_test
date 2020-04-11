package server

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"github.com/cloudrain21/learn_go_with_test/mytestapp/domain/server/store"
)

func TestGETPlayers(t *testing.T) {

	t.Run("in-memory store", func(t *testing.T) {
		store := store.NewInMemoryPlayerStore()
		server := NewPlayerServer(store)
		player := "Pepper"

		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assert.Equal(t, response.Code, http.StatusOK)
		//assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("stub store", func(t *testing.T) {
		store := store.NewStubPlayerStore()
		server := NewPlayerServer(store)
		player := "Floyd"

		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "2")
	})

	t.Run("not found", func(t *testing.T) {
		store := store.NewStubPlayerStore()
		server := NewPlayerServer(store)
		player := "dplee"

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusNotFound)
	})

	t.Run("it returns 200 on /league", func(t *testing.T) {
		sto := store.NewStubPlayerStore()
		server := NewPlayerServer(sto)

		request := httptest.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := store.League{}
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("decode error : %q : %v\n", response.Body, err)
		}

		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("league table", func(t *testing.T) {
		wantedLeague := store.League{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		sto := store.NewStubPlayerStore()
		server := NewPlayerServer(sto)

		for i := 0; i<32; i++ {
			server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest("Cleo"))
		}
		for i := 0; i<20; i++ {
			server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest("Chris"))
		}
		for i := 0; i<14; i++ {
			server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest("Tiest"))
		}

		request := httptest.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assert.Equal(t, "application/json", response.Header().Get("content-type"))

		var got store.League
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("decode fail : %v\n", err)
		}

		assertStatus(t, response.Code, http.StatusOK)

		if !reflect.DeepEqual(wantedLeague, got) {
			t.Errorf("want : %v got : %v\n", wantedLeague, got)
		}
	})

	t.Run("in-memory league", func(t *testing.T) {
		sto := store.NewInMemoryPlayerStore()
		server := NewPlayerServer(sto)

		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest("dplee"))

		request := httptest.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := store.League{
			{"dplee", 1},
		}
		var got store.League
		err := json.NewDecoder(response.Body).Decode(&got)
		assert.Equal(t, err, nil)

		assert.Equal(t, want, got)
	})
}

func newPostWinRequest(name string) *http.Request {
	request := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func newGetScoreRequest(name string) *http.Request {
	request := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want : %v got : %v\n", want, got)
	}
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Fatalf("want : %d got : %d\n", want, got)
	}
}

func CreateTempFile(t *testing.T, data string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "data_tmpfile")
	assert.Equal(t, nil, err)

	tmpfile.Write([]byte(data))

	removeFunc := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFunc
}

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database, removeDatabase := CreateTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		defer removeDatabase()

		sto, _ := store.NewFileSystemPlayerStore(database)

		got := sto.GetLeagueTable()

		want := store.League{
			{"Chris", 33},
			{"Cleo", 10},
		}
		got = sto.GetLeagueTable()

		assert.Equal(t, want, got)
	})

	t.Run("get player score", func(t *testing.T) {
		database, removeFunc := CreateTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)

		defer removeFunc()

		store, _ := store.NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Chris")
		want := 33

		assert.Equal(t, want, got)
	})

	t.Run("poset player score", func(t *testing.T) {
		database, removeFunc := CreateTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)

		defer removeFunc()

		store, _ := store.NewFileSystemPlayerStore(database)

		store.PostPlayerScore("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		assert.Equal(t, want, got)
	})

	t.Run("post player score", func(t *testing.T) {
		database, removeFunc := CreateTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)

		defer removeFunc()

		store, err := store.NewFileSystemPlayerStore(database)
		assert.Equal(t, nil, err)

		store.PostPlayerScore("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		assert.Equal(t, want, got)
	})

	t.Run("json parsing error", func(t *testing.T) {
		database, removeFunc := CreateTempFile(t, `[
        {"Name": "Cleo", "Wins": 10,
        {"Name": "Chris", "Wins": 33}]`)

		defer removeFunc()

		_, err := store.NewFileSystemPlayerStore(database)
		assert.Equal(t, store.ErrJsonParse, err)
	})

	t.Run("sorting", func(t *testing.T) {
		database, removeDatabase := CreateTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		defer removeDatabase()

		sto, _ := store.NewFileSystemPlayerStore(database)

		got := sto.GetLeagueTable()

		want := store.League{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assert.Equal(t, want, got)
	})
}
