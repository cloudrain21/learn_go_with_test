package main

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
)

func TestGETPlayers(t *testing.T) {

	t.Run("in-memory store", func(t *testing.T) {
		store := NewInMemoryPlayerStore()
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
		store := NewStubPlayerStore()
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
		store := NewStubPlayerStore()
		server := NewPlayerServer(store)
		player := "dplee"

		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		assertStatus(t, response.Code, http.StatusNotFound)
	})

	t.Run("it returns 200 on /league", func(t *testing.T) {
		store := NewStubPlayerStore()
		server := NewPlayerServer(store)

		request := httptest.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got League
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("decode error : %q : %v\n", response.Body, err)
		}

		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("league table", func(t *testing.T) {
		wantedLeague := League{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		store := StubPlayerStore{nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := httptest.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assert.Equal(t, "application/json", response.Header().Get("content-type"))

		var got League
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
		store := NewInMemoryPlayerStore()
		server := NewPlayerServer(store)

		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest("dplee"))

		request := httptest.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		want := League{
			{"dplee", 1},
		}
		var got League
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

		store := FileSystemPlayerStore{database}

		got := store.GetLeagueTable()

		want := League{
			{"Cleo", 10},
			{"Chris", 33},
		}
		got = store.GetLeagueTable()

		assert.Equal(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, removeFunc := CreateTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)

		defer removeFunc()

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Chris")
		want := 33

		assert.Equal(t, want, got)
	})

	t.Run("poset player score", func(t *testing.T) {
		database, removeFunc := CreateTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)

		defer removeFunc()

		store := FileSystemPlayerStore{database}

		store.PostPlayerScore("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		assert.Equal(t, want, got)
	})

	t.Run("poset player score", func(t *testing.T) {
		database, removeFunc := CreateTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)

		defer removeFunc()

		store := FileSystemPlayerStore{database}

		store.PostPlayerScore("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		assert.Equal(t, want, got)
	})
}
