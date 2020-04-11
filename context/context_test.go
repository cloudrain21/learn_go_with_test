package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response string
}

func (s *StubStore) Fetch(c context.Context) (string, error) {
	return s.response, nil
}

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, d := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("signal caught1")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(d)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case d := <-data:
		return d, nil
	}
}

func TestServer(t *testing.T) {
	t.Run("stub store", func(t *testing.T) {
		data := "hello, world"
		svr := Server(&StubStore{data})

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("want : %s, got : %s\n", data, response.Body.String())
		}
	})

	t.Run("context test1", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() != data {
			t.Errorf("want : %s, got : %s\n", data, response.Body.String())
		}
	})
}
