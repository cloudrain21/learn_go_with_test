package concurrency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {
	urlA := "www.daum.net"
	urlB := "www.naver.com"

	got := WebsiteRacer(urlA, urlB)
	want := urlA
	if got != want {
		t.Errorf("got : %s, want : %s\n", got, want)
	}
}

func TestRacer(t *testing.T) {
	t.Run("basic test", func(t *testing.T) {
		slowServer := makeDelayedServer(200 * time.Millisecond)
		fastServer := makeDelayedServer(100 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, _ := Racer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("got : %s, want : %s\n", got, want)
		}
	})

	//t.Run("timeout test", func(t *testing.T) {
	//	server1 := makeDelayedServer(5)
	//	server2 := makeDelayedServer(6)
	//
	//	defer server1.Close()
	//	defer server2.Close()
	//
	//	_, err := Racer(server1.URL, server2.URL)
	//	if err != nil {
	//		t.Error("timeout from server")
	//	}
	//})

	t.Run("configurable timeout test", func(t *testing.T) {
		server := makeDelayedServer(300 * time.Millisecond)
		defer server.Close()

		serverURL := server.URL

		_, err := ConfigurableRacer(serverURL, serverURL, 100 * time.Millisecond)
		if err != nil {
			t.Error("configurable timeout")
		}
	})
}

func makeDelayedServer(givenDelay time.Duration) (srv *httptest.Server) {
	srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(givenDelay)
		w.WriteHeader(http.StatusOK)
	}))
	return
}