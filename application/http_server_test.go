package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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

		var got []Player
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("decode error : %q : %v\n", response.Body, err)
		}

		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("league table", func(t *testing.T) {
		wantedLeague := []Player{
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

		var got []Player
		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("decode fail : %v\n", err)
		}

		assertStatus(t, response.Code, http.StatusOK)

		if !reflect.DeepEqual(wantedLeague, got) {
			t.Errorf("want : %v got : %v\n", wantedLeague, got)
		}
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
